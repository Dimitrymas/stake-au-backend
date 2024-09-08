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
	"log"
)

func main() {
	db, cancel, err := utils.DatabaseConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	log.Println("Connected to database")
	defer cancel()

	userCollection := db.Collection("users")
	userRepository := userPkg.NewRepository(userCollection)
	userService := userPkg.NewService(userRepository)
	userCommonHandler := user.NewCommonHandler(userService)

	accountCollection := db.Collection("accounts")
	accountRepository := accountPkg.NewRepository(accountCollection)
	accountService := accountPkg.NewService(accountRepository)
	accountCommonHandler := account.NewCommonHandler(accountService, userService)

	activationCollection := db.Collection("activations")
	activationRepository := activationPkg.NewRepository(activationCollection)
	activationService := activationPkg.NewService(activationRepository)
	activationCommonHandler := activation.NewCommonHandler(activationService)

	promocodeCollection := db.Collection("promocodes")
	promocodeRepository := promoCodePkg.NewRepository(promocodeCollection)
	promocodeService := promoCodePkg.NewService(promocodeRepository, activationService)
	promocodeCommonHandler := promocode.NewCommonHandler(promocodeService)

	app := fiber.New()
	routes.Router(
		app,
		userCommonHandler,
		promocodeCommonHandler,
		activationCommonHandler,
		accountCommonHandler,
	)

	log.Fatal(app.Listen(":3000"))
}
