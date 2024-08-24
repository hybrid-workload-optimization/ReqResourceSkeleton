package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	ys "main/ystruct"
)

func ResourceHandler(ctx *gin.Context) {

	recvData, err := ctx.GetRawData()
	if err != nil {
		ctx.YAML(400, gin.H{
			"message": fmt.Sprintf("%+v", err),
		})
		return
	}

	time.Sleep(time.Millisecond * 20)

	recvStr := fmt.Sprintf("%s", recvData)
	log.Printf("[RecvMsg]: " + recvStr)

	// sample response
	ctx.YAML(http.StatusOK, ys.RespResource{
		Response: ys.Response{
			ID:   "eddieKim",
			Date: "2024-0101 14:25:59",
			Result: ys.Result{
				Cluster:          "123",
				Node:             "Node01",
				PriorityClass:    "criticalPriority",
				Priority:         "100000",
				PreemptionPolicy: "Naver",
			},
		},
	})
}

func FinalHandler(ctx *gin.Context) {

	recvData, err := ctx.GetRawData()
	if err != nil {
		ctx.YAML(400, gin.H{
			"message": fmt.Sprintf("%+v", err),
		})
		return
	}

	time.Sleep(time.Millisecond * 20)

	recvStr := fmt.Sprintf("%s", recvData)
	log.Printf("[RecvMsg]: " + recvStr)

	ctx.YAML(http.StatusOK, gin.H{
		"status": "test",
		"value":  "OK",
		"msg":    "",
	})
}
