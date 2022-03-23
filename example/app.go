package main

import (
	"github.com/ujunglangit-id/tiny-memcache/example/entity"
	"github.com/ujunglangit-id/tiny-memcache/pkg/lib"
	"log"
)

func main() {
	cache := lib.NewCacheContainer(false, "./files/cache_restore_data")
	//init cache data
	kapur := entity.ProductDetail{
		ProductName:  "kapur barus",
		ProductReady: true,
		ProductStock: 10,
		CategoryName: "alat tulis",
	}
	pulpen := entity.ProductDetail{
		ProductName:  "pulpen",
		ProductReady: true,
		ProductStock: 5,
		CategoryName: "alat tulis",
	}

	err := cache.AddStructByKey("kapur", kapur)
	if err != nil {
		log.Fatalf("error AddStructByKey kapur, %#v", err)
	}
	err = cache.AddStructByKey("pulpen", pulpen)
	if err != nil {
		log.Fatalf("error AddStructByKey pulpen, %#v", err)
	}

	//get cache data
	var infoKapur, infoPulpen entity.ProductDetail
	err = cache.GetStructByKey("kapur", &infoKapur)
	if err != nil {
		log.Fatalf("error GetStructByKey kapur, %#v", err)
	}

	err = cache.GetStructByKey("pulpen", &infoPulpen)
	if err != nil {
		log.Fatalf("error GetStructByKey pulpen, %#v", err)
	}

	log.Printf("kapur : %#v", infoKapur)
	log.Printf("pulpen : %#v", infoPulpen)
}
