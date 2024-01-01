package main

import (
	"fmt"
	"math/rand"
)

type IObjectSizeVisitor interface {
	VisitDatabase(databaseService *DatabaseService) int64
	VisitFileStorage(fileStorageService *FileStorageService) int64
	VisitComputeService(computeService *ComputeService) int64
}

type ObjectSizeVisitorImpl struct{}

func NewObjectSizeVisitor() IObjectSizeVisitor {
	return &ObjectSizeVisitorImpl{}
}

func (visitor *ObjectSizeVisitorImpl) VisitDatabase(databaseService *DatabaseService) int64 {
	if databaseService.Tables > 0 {
		return databaseService.CalculateSize()
	} else {
		return 0
	}
}
func (visitor *ObjectSizeVisitorImpl) VisitFileStorage(fileStorageService *FileStorageService) int64 {
	// Load file...
	// Get file's size...
	return 122425 // KB
}
func (visitor *ObjectSizeVisitorImpl) VisitComputeService(computeService *ComputeService) int64 {
	return computeService.CalculateProductSize()
}

type IService interface {
	Accept(visitor IObjectSizeVisitor) int64
}

type DatabaseService struct {
	Tables int
	Rows   int64
}

func (service *DatabaseService) CalculateSize() int64 {
	return (int64)(service.Tables) * service.Rows
}
func (service *DatabaseService) Accept(visitor IObjectSizeVisitor) int64 {
	return visitor.VisitDatabase(service)
}

type FileStorageService struct {
	FilesCount int
}

func (storage *FileStorageService) GetStorageSize() int64 {
	var size int64 = 0
	for i := 0; i < storage.FilesCount; i++ {
		size += rand.Int63()
	}
	return size
}
func (storage *FileStorageService) Accept(visitor IObjectSizeVisitor) int64 {
	return visitor.VisitFileStorage(storage)
}

type ComputeService struct {
	Price    int64
	Quantity int64
}

func (service *ComputeService) CalculateProductSize() int64 {
	return service.Price * service.Quantity
}
func (service *ComputeService) Accept(visitor IObjectSizeVisitor) int64 {
	return visitor.VisitComputeService(service)
}

func main() {
	visitor := NewObjectSizeVisitor()

	userServices := []IService{
		&DatabaseService{Tables: 5, Rows: 12},
		&FileStorageService{FilesCount: 13},
		&ComputeService{Price: 20, Quantity: 5},
	}

	var globalSize int64 = 0
	for i, service := range userServices {
		size := service.Accept(visitor)
		fmt.Printf("[Service %d] Size: %d\n", i, size)
		globalSize += size
	}
	fmt.Println("Global size: ", globalSize)
}
