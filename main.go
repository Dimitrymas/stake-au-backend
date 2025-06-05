package main

import (
	"backend/api/http/handlers/account"
	"backend/api/http/handlers/activation"
	"backend/api/http/handlers/promocode"
	"backend/api/http/handlers/user"
	"backend/api/http/routes"
	accountPkg "backend/api/pkg/account"
	activationPkg "backend/api/pkg/activation"
	promoCodePkg "backend/api/pkg/promocode"
	userPkg "backend/api/pkg/user"
	"backend/api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"log"
)

func main() {
	db, cancel, err := utils.DatabaseConnection()
	if err != nil {
		log.Fatalf("Database Connection Error %s", err)
	}
	log.Println("Connected to database")
	defer cancel()

	userCollection := db.Collection("users")
	userRepository := userPkg.NewRepository(userCollection)
	userService := userPkg.NewService(userRepository)
	userCommonHandler := user.NewCommonHandler(userService)

	activationCollection := db.Collection("activations")
	activationRepository := activationPkg.NewRepository(activationCollection)
	activationService := activationPkg.NewService(activationRepository)
	activationCommonHandler := activation.NewCommonHandler(activationService)

	promocodeCollection := db.Collection("promocodes")
	promocodeRepository := promoCodePkg.NewRepository(promocodeCollection)
	promocodeService := promoCodePkg.NewService(promocodeRepository, activationService)
	promocodeCommonHandler := promocode.NewCommonHandler(promocodeService)

	accountCollection := db.Collection("accounts")
	accountRepository := accountPkg.NewRepository(accountCollection)
	accountService := accountPkg.NewService(accountRepository, userService, promocodeService, activationService)
	accountCommonHandler := account.NewCommonHandler(accountService, userService)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(helmet.New())

	routes.Router(
		app,
		userCommonHandler,
		promocodeCommonHandler,
		activationCommonHandler,
		accountCommonHandler,
	)

	log.Fatal(app.Listen(":3000"))
}
