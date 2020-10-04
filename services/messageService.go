package services

import (
	"fmt"
	"quasarFire/models"
	"strings"
)

type MessageService struct{}

func(l MessageService) GetMessageFromTopSecret(topSecret models.TopSecret) (message string){

	baseMessageToCompleteArray := topSecret.Satellites[0].Message
	for j := 0; j < len(topSecret.Satellites); j +=1{
		for i := 0; i < len(baseMessageToCompleteArray); i += 1{
			if topSecret.Satellites[j].Message[i] != ""{
				baseMessageToCompleteArray[i] = topSecret.Satellites[j].Message[i]
			}

		}
	}
	result := strings.Join(baseMessageToCompleteArray, " ")

	fmt.Println(result)

	return result

}