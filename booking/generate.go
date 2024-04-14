package generate

//go:generate swagger generate client -q -f ./docs/swagger.json -c booking_client -m booking_client/dto
//go:generate mockery
