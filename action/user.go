package action

import (
	"go-grpc-app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"google.golang.org/grpc"
	"log"
	"go-grpc-app/pb"
	"context"
	"time"
)

var userService *service.UserService

func init() {
	userService = service.NewUserService()
}

func GetUsers(gContext *gin.Context) {

	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	userService := pb.NewUserServiceClient(conn)
	ctx, cannel := context.WithTimeout(context.Background(), time.Second)
	defer cannel()
	result, err := userService.GetUsers(ctx, &pb.UserRequest{})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}

	gContext.JSON(http.StatusOK, gin.H{
		"errno":   0,
		"message": "ok",
		"data":    result.Users,
	})

}
