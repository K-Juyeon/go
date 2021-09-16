package main

import (
	"fmt"
	"math/rand"
	"time" //seed 생성용 패키지
)

//난수 추출된 수의 소수판정 프로그램 v0.6
//소수는 1과 자기 자신외에는 나누어 떨어지지 않는 수(0과 1 제외)
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2 //0, 1 제외, 2 ~ 151 사이의 수
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ { //1과 number일때 loop돌지 않음
		if number%i == 0 {
			isPrime = false
			break //첫번째 약수가 발견되면 반복문 즉시 종료
		}
		//fmt.Print(i, " ")
	}
	if isPrime { //비교 연산자 제거
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}
