package service

import (
	"context"
	"testing"
	"time"

	"github.com/huytran2000-hcmus/bTaskee/pricing/internal/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetPrice(t *testing.T) {
	Convey("Subject: Calculate the price by datetime", t, func() {
		svc := NewPricing()
		Convey("Given working details of a task", func() {
			req := model.CalculatePriceRequest{
				HouseType:    model.TwoRoom,
				ServiceTypes: []model.ServiceType{model.Cleaning, model.ChildCare},
				FromTime:     time.Date(2024, 4, 8, 15, 0, 0, 0, time.UTC),
				ToTime:       time.Date(2024, 4, 8, 17, 0, 0, 0, time.UTC),
			}
			Convey("When begin time and end time are on the same day", func() {
				Convey("When it is 2 hours of working", func() {

					Convey("When it is a house with 2 rooms", func() {
						req.HouseType = model.TwoRoom

						Convey("When it is on a weekday", func() {
							req.FromTime = time.Date(2024, 4, 8, 15, 0, 0, 0, time.UTC)
							req.ToTime = req.FromTime.Add(2 * time.Hour)

							Convey("The price is 200", func() {
								price, err := svc.GetPrice(context.Background(), req)
								So(err, ShouldEqual, nil)
								So(price, ShouldEqual, 200)
							})
						})

						Convey("When it is on the weekend", func() {
							req.FromTime = time.Date(2024, 4, 14, 15, 0, 0, 0, time.UTC)
							req.ToTime = req.FromTime.Add(2 * time.Hour)

							Convey("The price is 300", func() {
								price, err := svc.GetPrice(context.Background(), req)
								So(err, ShouldEqual, nil)
								So(price, ShouldEqual, 300)
							})

						})

					})

					Convey("When it is a house with 3 rooms", func() {
						req.HouseType = model.ThreeRoom
						Convey("When it is on a weekday", func() {
							req.FromTime = time.Date(2024, 4, 9, 15, 0, 0, 0, time.UTC)
							req.ToTime = req.FromTime.Add(2 * time.Hour)
							Convey("The price is 300", func() {
								price, err := svc.GetPrice(context.Background(), req)
								So(err, ShouldEqual, nil)
								So(price, ShouldEqual, 300)
							})

						})
						Convey("When it is on the weekend", func() {
							req.FromTime = time.Date(2024, 4, 14, 15, 0, 0, 0, time.UTC)
							req.ToTime = req.FromTime.Add(2 * time.Hour)

							Convey("The price is 450", func() {
								price, err := svc.GetPrice(context.Background(), req)
								So(err, ShouldEqual, nil)
								So(price, ShouldEqual, 450)
							})
						})
					})

					Convey("When it is a house with 4 rooms", func() {
						req.HouseType = model.FourRoom

						Convey("When it is on a weekday", func() {
							req.FromTime = time.Date(2024, 4, 8, 15, 0, 0, 0, time.UTC)
							req.ToTime = req.FromTime.Add(2 * time.Hour)

							Convey("The price is 400", func() {
								price, err := svc.GetPrice(context.Background(), req)
								So(err, ShouldEqual, nil)
								So(price, ShouldEqual, 400)
							})

						})

						Convey("When it is on the weekend", func() {
							req.FromTime = time.Date(2024, 4, 14, 15, 0, 0, 0, time.UTC)
							req.ToTime = req.FromTime.Add(2 * time.Hour)

							Convey("The price is 600", func() {
								price, err := svc.GetPrice(context.Background(), req)
								So(err, ShouldEqual, nil)
								So(price, ShouldEqual, 600)
							})

						})

					})
				})
			})

		})
	})
}

func TestGetPrice_Error(t *testing.T) {
	Convey("Subject: Calculate the price by datetime", t, func() {
		svc := NewPricing()
		Convey("Given working details of a task", func() {
			req := model.CalculatePriceRequest{
				HouseType:    model.TwoRoom,
				ServiceTypes: []model.ServiceType{model.Cleaning, model.ChildCare},
				FromTime:     time.Date(2024, 4, 8, 15, 0, 0, 0, time.UTC),
				ToTime:       time.Date(2024, 4, 8, 17, 0, 0, 0, time.UTC),
			}

			Convey("When begin time and end time are on the same day", func() {
				Convey("When it is 2 hours of working", func() {
					Convey("But it is unknown house type", func() {
						req.HouseType = "zillion-room"

						Convey("Return an error indicating unknown house type", func(c C) {
							price, err := svc.GetPrice(context.Background(), req)
							So(price, ShouldEqual, 0)
							So(err, ShouldWrap, ErrHouseTypeNotFound)
						})

					})
				})

				Convey("But it is less than 2 hours of working", func() {
					req.FromTime = time.Date(2024, 4, 14, 15, 0, 0, 0, time.UTC)
					req.ToTime = req.FromTime.Add(1 * time.Hour)

					Convey("Return an error indicating less than minimum working hours", func() {
						price, err := svc.GetPrice(context.Background(), req)
						So(price, ShouldEqual, 0)
						So(err, ShouldWrap, ErrWorkingLessThanMinimum)
					})
				})

				Convey("But it is more than 4 hours of working", func() {
					req.FromTime = time.Date(2024, 4, 14, 15, 0, 0, 0, time.UTC)
					req.ToTime = req.FromTime.Add(5 * time.Hour)

					Convey("Return an error indicating less than minimum working hours", func() {
						price, err := svc.GetPrice(context.Background(), req)
						So(price, ShouldEqual, 0)
						So(err, ShouldWrap, ErrWorkingMoreThanMaximum)
					})
				})
			})

			Convey("But the begin time and end time are on different days", func() {
				req.FromTime = time.Date(2024, 4, 14, 15, 0, 0, 0, time.UTC)
				req.ToTime = req.FromTime.Add(24 * time.Hour)
				Convey("Return an error", func() {
					price, err := svc.GetPrice(context.Background(), req)
					So(price, ShouldEqual, 0)
					So(err, ShouldWrap, ErrWorkingsNotInTheSameDate)
				})

			})

		})
	})
}
