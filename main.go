package main

import (
	"context"
	"fmt"

	"github.com/thara/ogen-playground/api"
)

func main() {

}

type handler struct {
	api.UnimplementedHandler
}

func (h *handler) Read(ctx context.Context, params api.ReadParams) (r *api.ThePet, err error) {
	return &api.ThePet{
		Pet: api.Pet{
			Name:   "",
			Weight: api.NewOptFloat32(10),
			Kind:   "x",
			// wrong: not found Cat.meow, Dog.dark
		},
		Pet2: api.Pet2{
			Name:   "",
			Weight: api.NewOptFloat32(10),
			ID:     "x",
			// wrong: not found Cat.meow, Dog.dark
		},

		// Pet3: NG(ogen not supported)

		// oneof

		// OK
		Pet4: api.Pet4{
			Type: api.Cat3Pet4,
			Cat3: api.Cat3{
				Meow: api.NewOptInt32(10),
			},
			Dog3: api.Dog3{
				Bark: api.NewOptString("22"),
			},
		},
		// OK
		Pet5: api.Pet5{
			OneOf: api.NewCat5Pet5Sum(api.Cat5{
				Kind: api.Cat5KindCat,
				Meow: api.NewOptInt32(10),
			}),
		},
		// OK
		Pet6: api.Pet6{
			OneOf: api.NewPet6CatPet6Sum(api.Pet6Cat{
				DataKind: api.Pet6CatDataKindCat,
				Data: api.Cat6{
					Name: "string",
					Meow: false,
				},
			}),
		},
		// OK
		Pet7: api.Pet7{
			OneOf: api.NewDog7Pet7Sum(api.Dog7{
				Name: "string",
				Bark: false,
			}),
		},
		// OK
		Pet8: api.Pet8{
			OneOf: api.NewDog8Pet8Sum(api.Dog8{
				Bark: api.NewOptString(""),
			}),
		},
	}, nil
}

func (h *handler) Create(ctx context.Context, req *api.CreateReq, params api.CreateParams) (r *api.ThePet, err error) {
	if req.Body.Pet.Kind == "cat" {
		//NG: missing field
	}
	if req.Body.Pet2.ID == "cat" {
		//NG: missing field
	}

	// Pet3: NG(ogen not supported)

	// oneof

	// OK
	switch req.Body.Pet4.Type {
	case api.Cat3Pet4:
		fmt.Println(req.Body.Pet4.Cat3)
	case api.Dog3Pet4:
		fmt.Println(req.Body.Pet4.Dog3)
	}

	// OK
	switch req.Body.Pet5.OneOf.Type {
	case api.Cat5Pet5Sum:
		fmt.Println(req.Body.Pet5.OneOf.Cat5)
	case api.Dog5Pet5Sum:
		fmt.Println(req.Body.Pet5.OneOf.Dog5)
	}

	// OK
	switch req.Body.Pet6.OneOf.Type {
	case api.Pet6CatPet6Sum:
		fmt.Println(req.Body.Pet6.OneOf.Pet6Cat)
	case api.Pet6DogPet6Sum:
		fmt.Println(req.Body.Pet6.OneOf.Pet6Dog)
	}

	// OK
	switch req.Body.Pet7.OneOf.Type {
	case api.Cat7Pet7Sum:
		fmt.Println(req.Body.Pet7.OneOf.Cat7)
	case api.Dog7Pet7Sum:
		fmt.Println(req.Body.Pet7.OneOf.Dog7)
	}

	// OK
	switch req.Body.Pet8.OneOf.Type {
	case api.Cat8Pet8Sum:
		fmt.Println(req.Body.Pet8.OneOf.Cat8)
	case api.Dog8Pet8Sum:
		fmt.Println(req.Body.Pet8.OneOf.Dog8)
	}

	return &api.ThePet{}, nil
}

func (h *handler) List(ctx context.Context, params api.ListParams) (r []api.ThePet, err error) {
	for _, t := range params.Tag {
		fmt.Println(t)
	}
	return r, nil
}
