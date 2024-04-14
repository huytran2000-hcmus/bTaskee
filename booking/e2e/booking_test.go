package e2e

import (
	"context"
	"fmt"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/huytran2000-hcmus/bTaskee/booking/booking_client"
	"github.com/huytran2000-hcmus/bTaskee/booking/booking_client/dto"
	"github.com/huytran2000-hcmus/bTaskee/booking/booking_client/task"
	"github.com/huytran2000-hcmus/bTaskee/booking/config"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/infra/mongodb"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/mocks/repository"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/model"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/restful"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"github.com/testcontainers/testcontainers-go"
	mongodbctn "github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testSuite struct {
	*httptest.Server
	client      *booking_client.BookingAPI
	dbContainer *mongodbctn.MongoDBContainer
	cfg         suiteCfg
	mongoDB     *mongo.Database
}

type suiteCfg struct {
	srvPort    int
	dbUsername string
	dbPassword string
	dbName     string
}

func TestCreateTask(t *testing.T) {
	ts := newTestSuite(t, suiteCfg{
		srvPort:    8800,
		dbUsername: "bTaskee",
		dbPassword: "secret",
		dbName:     "booking",
	})
	defer ts.Close()

	Convey("Subject: Create Task", t, func() {
		Convey("Given a task request", func() {
			customerAddress := "60 Bến Thành, Quận 1, TP.Hồ Chí Minh"
			customerEmail := "huy@gmail.com"
			customerName := "Huy"
			customerPhone := "+84948447524"
			note := "Be here ontime"
			taskerID := "1"
			taskerName := "Khương"
			taskerPhone := "+84948447525"
			taskerCMND := "888888888"
			taskerCCCD := "999999999999"
			taskerEmail := "khuong@gmail.com"
			fromTime := time.Date(2024, 4, 7, 9, 0, 0, 0, time.Local)
			toTime := time.Date(2024, 4, 7, 11, 0, 0, 0, time.Local)
			fromTimeStr := fromTime.Format(time.RFC3339)
			toTimeStr := toTime.Format(time.RFC3339)
			houseType := "two_room"
			body := &dto.ModelCreateTaskRequest{
				AssignedLocation: &dto.ModelLocation{
					Address: &customerAddress,
				},
				Customer: &dto.ModelCustomer{
					Email: &customerEmail,
					Name:  &customerName,
					Phone: &customerPhone,
				},
				Note: note,
				Tasker: &dto.ModelTasker{
					Email: &taskerEmail,
					ID:    &taskerID,
					Identification: &dto.ModelIdentification{
						Cccd: taskerCCCD,
						Cmnd: taskerCMND,
					},
					Name:  &taskerName,
					Phone: &taskerPhone,
				},
				WorkingDetails: &dto.ModelWorkingDetails{
					FromTime:     &fromTimeStr,
					ServiceTypes: []dto.ModelServiceType{dto.ModelServiceTypeCleaning},
					ToTime:       &toTimeStr,
					HouseType:    &houseType,
				},
			}
			param := task.NewPostAPIV1TasksParams().WithDefaults().WithRequest(body)

			Convey("When the task request is submitted successfully", func() {
				resp, err := ts.client.Task.PostAPIV1Tasks(param)
				So(err, ShouldEqual, nil)
				payload := resp.Payload
				So(payload, ShouldNotBeNil)

				idStr := payload.Data
				id, err := uuid.Parse(idStr)
				So(err, ShouldEqual, nil)

				Convey("When get a task in the database with the id equal response's id", func() {
					task := ts.getTask(id)

					Convey("It should have the same informations", func() {
						So(task.ID, ShouldEqual, id)
						So(task.Customer.Name, ShouldEqual, customerName)
						So(task.Customer.Email, ShouldEqual, customerEmail)
						So(task.Customer.Phone, ShouldEqual, customerPhone)
						So(task.AssignedLocation.Address, ShouldEqual, customerAddress)
						So(task.WorkingDetails.From, ShouldEqual, fromTime)
						So(task.WorkingDetails.To, ShouldEqual, toTime)
						So(task.WorkingDetails.HouseType, ShouldEqual, houseType)
						So(task.WorkingDetails.ServiceTypes, ShouldResemble, []string{string(dto.ModelServiceTypeCleaning)})
						So(task.Note, ShouldEqual, note)
						So(task.Tasker.ID, ShouldEqual, taskerID)
						So(task.Tasker.Name, ShouldEqual, taskerName)
						So(task.Tasker.Email, ShouldEqual, taskerEmail)
						So(task.Tasker.Phone, ShouldEqual, taskerPhone)
						So(task.Tasker.Identification.CCCD, ShouldEqual, taskerCCCD)
						So(task.Tasker.Identification.CMND, ShouldEqual, taskerCMND)
					})
				})
			})

		})

	})
}

func newTestSuite(t *testing.T, cfg suiteCfg) *testSuite {
	ts := &testSuite{
		cfg: cfg,
	}

	ts.setupDB(t)
	ts.setupEnv(t)
	err := config.SetEnv()
	if err != nil {
		t.Fatal(err)
	}

	ts.setupDB(t)

	sendRepo := repository.NewMockSendRepository(t)
	sendRepo.EXPECT().SendTask(mock.Anything, mock.Anything).Return(model.SendTaskResponse{
		MessageIDs: []string{"1"},
	}, nil)
	pricingRepo := repository.NewMockPricingRepository(t)
	pricingRepo.EXPECT().GetPricing(mock.Anything, mock.Anything).Return(model.CalculatePriceResponse{
		Total: 100,
	}, nil)
	app, err := restful.New(pricingRepo, sendRepo)
	if err != nil {
		t.Fatal(err)
	}

	ts.Server = httptest.NewServer(app.Routes())
	ts.setupClient(t)
	ts.setupMongoClient(t)

	return ts
}

func (ts *testSuite) setupDB(t *testing.T) {
	ctx := context.Background()
	var err error
	ts.dbContainer, err = mongodbctn.RunContainer(
		ctx,
		testcontainers.WithImage("mongo:7.0.8"),
		testcontainers.WithEnv(map[string]string{
			"MONGO_INITDB_ROOT_USERNAME": ts.cfg.dbUsername,
			"MONGO_INITDB_ROOT_PASSWORD": ts.cfg.dbPassword,
		}),
	)
	if err != nil {
		t.Fatalf("failed to start mongo container: %s", err)
	}
}

func (ts *testSuite) setupClient(t *testing.T) {
	uri, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	transport := client.New(uri.Host, "/", []string{uri.Scheme})
	ts.client = booking_client.New(transport, strfmt.Default)
}

func (ts *testSuite) setupEnv(t *testing.T) {
	cfg := ts.cfg
	dbAddress, err := ts.dbContainer.Endpoint(context.Background(), "")
	if err != nil {
		t.Fatal(err)
	}

	os.Setenv("DB_URI", fmt.Sprintf("mongodb://%s:%s@%s", cfg.dbUsername, cfg.dbPassword, dbAddress))
}

func (ts *testSuite) setupMongoClient(t *testing.T) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opt := options.Client().ApplyURI(config.LoadEnv().DBURI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		t.Fatal(err)
	}

	var result bson.M
	db := client.Database(config.LoadEnv().DBName)
	err = db.RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}

	ts.mongoDB = db
}

func (ts *testSuite) getTask(id uuid.UUID) *mongodb.Task {
	task := &mongodb.Task{}
	filter := bson.D{{"_id", id}}
	ctx := context.Background()
	coll := ts.mongoDB.Collection("tasks")
	err := coll.FindOne(ctx, filter).Decode(task)
	So(err, ShouldEqual, nil)

	return task
}
