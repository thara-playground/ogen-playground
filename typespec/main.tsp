import "@typespec/http";
import "@typespec/openapi3";
using OpenAPI;

using Http;
@service(#{ title: "Service" })
namespace DemoService;

@error
model Error {
  code: int32;
  message: string;
}

@get
op read(@path id: string): ThePet | Error;

@post
op create(@path id: string, body: ThePet) : {
  @statusCode statusCode: 200;
  @body pet: ThePet;
};

model ThePet {
  pet: Pet;
  pet2: Pet2;
  /* pet3: Pet3; */
  pet4: Pet4;
  pet5: Pet5;
  pet6: Pet6;
  pet7: Pet7;
  pet8: Pet8;
}

@discriminator("kind")
model Pet {
  name: string;
  weight?: float32;
}
model Cat extends Pet {
  kind: "cat";
  meow: int32;
}
model Dog extends Pet {
  kind: "dog";
  bark: string;
}


@discriminator("id")
model Pet2 {
  name: string;
  weight?: float32;
}
model Cat2 extends Pet2 {
  id: "cat";
  meow?: int32;
}
model Dog2 extends Pet2 {
  id: "dog";
  bark?: string;
}

union Pet3 {
  cat: Cat3,
  dog: Dog3,
}

model Cat3 {
  meow?: int32;
}

model Dog3 {
  bark?: string;
}

@oneOf
union Pet4 {
  cat: Cat3,
  dog: Dog3,
}

@discriminated(#{ envelope: "none" })
union Pet5 {
  cat: Cat5,
  dog: Dog5,
}
model Cat5 {
  kind: "cat";
  meow?: int32;
}
model Dog5 {
  kind: "dog";
  bark?: string;
}

@discriminated(#{discriminatorPropertyName: "dataKind", envelopePropertyName: "data"})
union Pet6{ cat: Cat6, dog: Dog6 }

model Cat6 { name: string, meow: boolean }
model Dog6 { name: string, bark: boolean }

@discriminated(#{ envelope: "none", discriminatorPropertyName: "id" })
union Pet7{ `cat-cat-cat`: Cat7, `dog-dog-dog`: Dog7 }

model Cat7 { id: "cat-cat-cat", name: string, meow: boolean }
model Dog7 { id: "dog-dog-dog", name: string, bark: boolean }


@discriminated(#{ envelope: "none" })
union Pet8 {
  cat: Cat8,
  dog: Dog8,
}
model Cat8 {
  meow?: int32;
}
model Dog8 {
  bark?: string;
}
