package repository

type HelloWorldRepository interface {
	GetHelloWorld() string
}

type helloWorldRepository struct{}

func NewHelloWorldRepository() HelloWorldRepository {
	return &helloWorldRepository{}
}

func (r *helloWorldRepository) GetHelloWorld() string {
	// This will use the singleton instance
	//client, err := db.GetClient()
	//if err != nil {
	//	return nil, err
	//}
	//var version schema.AdonisSchema
	//client.Database().Model(&version).First(&version) // Use db variable to avoid unused variable warning
	//
	//fmt.Println(version)

	return "Hello, world!"
}
