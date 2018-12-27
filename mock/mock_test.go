package mock

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestMockRepository(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)

	mockRepo.EXPECT().Get(gomock.Any()).Return(User{Id: 1, Name: "tom", Age: 20}).AnyTimes()
	//mockRepo.EXPECT().Get(gomock.Any()).Return(User{Id: 2, Name: "jerry", Age: 30}).Times(1)
	//mockRepo.EXPECT().Get(gomock.Any()).Return(User{Id: 3, Name: "john", Age: 25}).MaxTimes(5)

	fmt.Println(mockRepo.Get(0)) // {1 tom 20}
	fmt.Println(mockRepo.Get(0)) // {2 jerry 30}
	fmt.Println(mockRepo.Get(0)) // {3 john 25}
	fmt.Println(mockRepo.Get(0))
	fmt.Println(mockRepo.Get(0))
	fmt.Println(mockRepo.Get(0))
}
