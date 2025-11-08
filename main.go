package main

import (
	"context"

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

		// oneof
		Pet4: api.Pet4{
			Type: api.Cat3Pet4,
			// OK
			Cat3: api.Cat3{
				Meow: api.NewOptInt32(10),
			},
			Dog3: api.Dog3{
				Bark: api.NewOptString("22"),
			},
		},
		Pet5: api.Pet5{
			OneOf: api.NewCat5Pet5Sum(api.Cat5{
				Kind: api.Cat5KindCat,
				Meow: api.NewOptInt32(10),
			}),
		},
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
		Pet8: api.Pet8{
			OneOf: api.NewDog8Pet8Sum(api.Dog8{
				Bark: api.NewOptString(""),
			}),
		},
	}, nil
}

func (h *handler) Create(ctx context.Context, req *api.CreateReq, params api.CreateParams) (r *api.ThePet, err error) {
	return &api.ThePet{}, nil
}
