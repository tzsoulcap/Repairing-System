package main

import (
	"github.com/tzsoulcap/ui-sa/controller"
	"github.com/tzsoulcap/ui-sa/entity"
	"github.com/tzsoulcap/ui-sa/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())
	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{ //****************Tenant
			//employer
			protected.GET("/employers", controller.ListEmployer)
			protected.GET("/employer/:id", controller.GetEmployer)
			protected.POST("/employers", controller.Createmployer)
			protected.PATCH("/employers", controller.UpdatEmployers)
			protected.DELETE("/employers/:id", controller.DeleteEmployers)

			//Prefix
			protected.GET("/prefixes", controller.ListPrefix)
			protected.GET("/prefixes/:id", controller.GetPrefix)
			protected.POST("/prefixes", controller.CreatePrefix)
			protected.PATCH("/prefixes", controller.UpdatePrefix)
			protected.POST("/prefixes/:id", controller.UpdatePrefixByID)
			protected.DELETE("/prefixes/:id", controller.DeletePrefix)

			//Career
			protected.GET("/careers", controller.ListCareer)
			protected.GET("/careers/:id", controller.GetCareer)
			protected.POST("/careers", controller.CreateCareer)
			protected.PATCH("/careers", controller.UpdateCareer)
			protected.DELETE("/careers/:id", controller.DeleteCareer)

			//Gender
			protected.GET("/genders", controller.ListGender)
			protected.GET("/genders/:id", controller.GetGender)
			protected.POST("/genders", controller.CreateGender)
			protected.PATCH("/genders", controller.UpdateGender)
			protected.DELETE("/genders/:id", controller.DeleteGender)

			//Canton
			protected.GET("/cantons", controller.ListCanton)
			protected.GET("/cantons/:id", controller.GetCanton)
			protected.POST("/cantons", controller.CreateCanton)
			protected.PATCH("/cantons", controller.UpdateCanton)
			protected.DELETE("/cantons/:id", controller.DeleteCanton)

			//Prefecture
			protected.GET("/prefectures", controller.ListPrefecture)
			protected.GET("/prefectures/:id", controller.GetPrefecturebyproviceid)
			protected.POST("/prefectures", controller.CreatePrefecture)
			protected.PATCH("/prefectures", controller.UpdatePrefecture)
			protected.DELETE("/prefectures/:id", controller.DeletePrefecture)

			//Province
			protected.GET("/provinces", controller.ListProvince)
			protected.GET("/provinces/:id", controller.GetProvince)
			protected.POST("/provinces", controller.CreateProvince)
			protected.PATCH("/provinces", controller.UpdateProvince)
			protected.DELETE("/provinces/:id", controller.DeleteProvince)

			//Tenant
			protected.GET("/tenants", controller.ListTenant)
			protected.GET("/tenant/:id", controller.GetTenant)
			protected.POST("/tenants", controller.CreateTenant)
			protected.PATCH("/tenants", controller.UpdateTenant)
			protected.DELETE("/tenants/:id", controller.DeleteTenant)

			//****************Rental
			// room Routes
			protected.GET("/rooms", controller.ListRoom)
			protected.GET("/room", controller.ListRooms_RoomReturn)
			protected.GET("/room/:id", controller.GetRoom)
			//Repair
			protected.GET("/rooms/:id", controller.GetRoomofRepairandroomreturn)
			protected.POST("/rooms", controller.CreateRoom)
			protected.PATCH("/rooms", controller.UpdateRoom)
			protected.DELETE("/rooms/:id", controller.DeleteRoom)

			// roomstate Routes
			protected.GET("/roomstates", controller.ListRoomstate)
			protected.GET("/roomstate/:id", controller.GetRoomstate)
			protected.POST("/roomstates", controller.CreateRoomstate)
			protected.PATCH("/roomstates", controller.UpdateRoomstate)
			protected.DELETE("/roomstates/:id", controller.DeleteRoomstate)

			// roomtype Routes
			protected.GET("/roomtypes", controller.ListRoomtype)
			protected.GET("/roomtype/:id", controller.GetRoomtype)
			protected.POST("/roomtypes", controller.CreateRoomtype)
			protected.PATCH("/roomtypes", controller.UpdateRoomtype)
			protected.DELETE("/roomtypes/:id", controller.DeleteRoomtype)

			// rentalstate Routes
			protected.GET("/rentalstates", controller.ListRentalstate)
			protected.GET("/rentalstate/:id", controller.GetRentalstate)
			protected.POST("/rentalstates", controller.CreateRentalstate)
			protected.PATCH("/rentalstates", controller.UpdateRentalstate)
			protected.DELETE("/rentalstates/:id", controller.DeleteRentalstate)

			// rental Routes
			protected.GET("/rentals", controller.ListRental)
			protected.GET("/rental/:id", controller.GetRental)
			//Repair
			protected.GET("/rentals/:id", controller.GetRentalofRepair)
			//Roomretrun
			protected.GET("/rentals/rooms/:id", controller.GetRentalofRoomretrun)
			protected.POST("/rentals", controller.CreateRental)
			protected.PATCH("/rentals", controller.UpdateRental)
			protected.DELETE("/rentals/:id", controller.DeleteRental)

			//****************Rental
			// Equipment Routes
			protected.GET("/Equipments", controller.ListEquipments)
			protected.GET("/Equipment/:id", controller.GetEquipment)
			protected.POST("/Equipments", controller.CreateEquipment)
			protected.PATCH("/Equipments", controller.UpdateEquipment)
			protected.DELETE("/Equipments/:id", controller.DeleteEquipment)

			// Room Equipment Routes
			protected.GET("/RoomEquipments/room/:number/type/:typeid", controller.ListRoomEquipments)
			protected.GET("/RoomEquipment/:id", controller.GetRoomEquipment)
			protected.POST("/RoomEquipments", controller.CreateRoomEquipment)
			protected.PATCH("/RoomEquipments", controller.UpdateRoomEquipment)
			protected.DELETE("/RoomEquipments/:id", controller.DeleteRoomEquipment)

			// Type Equipment Routes
			protected.GET("/TypeEquipments", controller.ListTypeEquipments)
			protected.GET("/TypeEquipment/:id", controller.GetTypeEquipment)
			protected.POST("/TypeEquipments", controller.CreateTypeEquipment)
			protected.PATCH("/TypeEquipments", controller.UpdateTypeEquipment)
			protected.DELETE("/TypeEquipments/:id", controller.DeleteTypeEquipment)

			// Repair Routes
			protected.GET("/Repairs", controller.ListRepairs)
			protected.GET("/Repair/:id", controller.GetRepair)
			protected.POST("/Repairs", controller.CreateRepair)
			protected.PATCH("/Repairs", controller.UpdateRepair)
			protected.DELETE("/Repairs/:id", controller.DeleteRepair)
			//****************RoomReturn
			// ReturnRoom Routes
			protected.GET("/return_rooms", controller.ListReturnRooms)
			protected.GET("/returnroom/:id", controller.GetReturnRoom)
			protected.POST("/return_rooms", controller.CreateReturnRoom)
			protected.PATCH("/return_rooms", controller.UpdateReturnRoom)
			protected.DELETE("/returnrooms/:id", controller.DeleteReturnRoom)

			//****************Promotion
			// Festival Routes

			protected.GET("/fests", controller.ListFests)
			protected.GET("/fest/:id", controller.GetFest)
			protected.POST("/fests", controller.CreateFest)
			protected.PATCH("/fests", controller.UpdateFest)
			protected.DELETE("/fests/:id", controller.DeleteFest)

			// Info Routes

			protected.GET("/infos", controller.ListInfo)
			protected.GET("/info/:id", controller.GetInfo)
			protected.POST("/infos", controller.CreateInfo)
			protected.PATCH("/infos", controller.UpdateInfo)
			protected.DELETE("/infos/:id", controller.DeleteInfo)

			// Type Routes

			protected.GET("/ptypes", controller.ListPtype)
			protected.GET("/ptype/:id", controller.GetPtype)
			protected.POST("/ptypes", controller.CreatePType)
			protected.PATCH("/ptypes", controller.UpdatePtype)
			protected.DELETE("/ptypes/:id", controller.DeletePtype)

			// Active Routes

			protected.GET("/actives", controller.ListActives)
			protected.GET("/active/:id", controller.GetActive)
			protected.POST("/actives", controller.CreateActive)
			protected.PATCH("/actives", controller.UpdateActive)
			protected.DELETE("/actives/:id", controller.DeleteActive)

			//****************Payment
			protected.GET("/payments", controller.ListPayments)
			protected.GET("/payment/:id", controller.GetPayment)
			protected.POST("/payments", controller.CreatePayment)
			protected.PATCH("/payments", controller.UpdatePayment)
			protected.DELETE("/payments/:id", controller.DeletePayment)
		}
	}

	r.POST("/login", controller.Login)
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
