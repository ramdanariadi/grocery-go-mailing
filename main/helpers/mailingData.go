package helpers

type Cc struct {
	Address string `bson:"address"`
	Name    string `bson:"name"`
}

type MailingData struct {
	To      []string `bson:"to"`
	Cc      Cc       `bson:"cc"`
	Subject []string `bson:"subject"`
	Body    string   `bson:"body"`
}
