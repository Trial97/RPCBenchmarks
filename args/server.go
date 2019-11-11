package args

import (
	"context"
	"fmt"
	"log"
	"os"
)

const (
	MyServerSum = "MyServer.Sum"
)

//MyServer used to sum the numbers
type MyServer struct{}

//Sum adds numbers
func (b *MyServer) Sum(a *Args, r *int) error {
	*r = a.A + a.B
	return nil
}

//WriteNumber write number to file;
func (b *MyServer) WriteNumber(a Argsw, reply *bool) error {
	fo, err := os.Create(a.F)
	if err != nil {
		log.Println("WriteNumberError:", err)
		return err
	}
	defer fo.Close()
	fmt.Fprintf(fo, "%d", a.A)
	*reply = true
	return nil
}

//ReadNumber reads number from file
func (b *MyServer) ReadNumber(f string, r *int) error {
	fd, err := os.Open(f)
	if err != nil {
		log.Println("ReadNumberError:", err)
		return err
	}
	fmt.Fscanf(fd, "%d\n", r)
	return nil
}

//MyServer2 used to sum the numbers for gorpc server implementation
type MyServer2 struct{}

//Sum adds numbers
func (b *MyServer2) Sum(a *Args) (int, error) {
	return a.A + a.B, nil
}

//MyServer used to sum the numbers
type MyServer3 struct{}

//Sum adds numbers
func (b *MyServer3) Sum(_ context.Context, a *Args, r *int) error {
	*r = a.A + a.B
	return nil
}
