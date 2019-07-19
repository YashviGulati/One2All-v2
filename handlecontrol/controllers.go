package handlecontrol

import
(
	"net/http"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/sns"
	"os"
	"gorilla/mux"
	"flag"
)

const (
	AKID string = ""
	SECRET_KEY string = ""
)

func Home(w http.ResponseWriter, r *http.Request)(){
	if r.URL.Path != "/"{
		error(w,r, http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Welcome to Homepage")
}

func GetTopics(w http.ResponseWriter, r *http.Request)(){

	if r.URL.Path != "/topics"{
		error(w,r, http.StatusNotFound)
		return
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials(AKID, SECRET_KEY, ""),
	},)

	svc := sns.New(sess)

	result, err := svc.ListTopics(nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, t := range result.Topics {
		fmt.Fprintln(w,*t.TopicArn)
	}
	fmt.Println("Topics listed")
}

func error(w http.ResponseWriter, r *http.Request, err int) {
	w.WriteHeader(err)
	if err == http.StatusNotFound {
		fmt.Fprint(w, "Error 404")
	}
}

func GetSubByTopic(w http.ResponseWriter, r *http.Request){
	value:=mux.Vars(r)
	deltemp:="arn:aws:sns:ap-south-1:210721209503:"+value["name"]
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials(AKID, SECRET_KEY, ""),
	},)

	svc := sns.New(sess)
	result, err:= svc.ListSubscriptionsByTopic(&sns.ListSubscriptionsByTopicInput{TopicArn: aws.String(deltemp)})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Fprintln(w,result)
	fmt.Println("Subscriptions Listed")

}

func DeleteTopicByName(w http.ResponseWriter, r *http.Request){
	value:= mux.Vars(r)
	deltemp:="arn:aws:sns:ap-south-1:210721209503:"+value["name"]
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials(AKID, SECRET_KEY, ""),
	},)
	svc := sns.New(sess)
	result, err := svc.DeleteTopic(&sns.DeleteTopicInput{ TopicArn: aws.String(deltemp)})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Fprintln(w,result)
	fmt.Println("Deletion Successful : ", value["name"])
}

func CreateSub(w http.ResponseWriter, r *http.Request){

	value:= mux.Vars(r)
	subname:=value["sub"]
	topname:="arn:aws:sns:ap-south-1:210721209503:"+value["name"]

	emailpt := flag.String("e", subname, "The email address of the user subscribing to the topic")
	topicpt := flag.String("t",topname, "The ARN of the topic to which the user subscribes")
	flag.Parse()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials(AKID, SECRET_KEY, ""),
	},)
	svc := sns.New(sess)

	result, err := svc.Subscribe(&sns.SubscribeInput{
		Endpoint:              emailpt,
		Protocol:              aws.String("email"),
		ReturnSubscriptionArn: aws.Bool(true), // Return the ARN, even if user has yet to confirm
		TopicArn:              topicpt,
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(*result.SubscriptionArn)
	fmt.Fprintln(w,"Subscription Created:  "+*result.SubscriptionArn)
}

func SendMsg(w http.ResponseWriter, r *http.Request){

	value:= mux.Vars(r)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials(AKID, SECRET_KEY, ""),
	},)
	if err != nil {
		fmt.Println("NewSession error:", err)
		return
	}
	endpt := sns.New(sess)
	msg := &sns.PublishInput{
		Message:  aws.String(value["msg"]),
		TopicArn: aws.String("arn:aws:sns:ap-south-1:210721209503:"+value["name"]),
	}

	result, err := endpt.Publish(msg)
	if err != nil {
		fmt.Println("Publish error:", err)
		return
	}
	fmt.Fprintln(w,result.String())
	fmt.Println("Message Sent")
}

func CreateTopic(w http.ResponseWriter, r *http.Request) {
	value := mux.Vars(r)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials(AKID, SECRET_KEY, ""),
	}, )
	svc := sns.New(sess)
	alltopics, err := svc.ListTopics(nil)
	deltemp := "arn:aws:sns:ap-south-1:210721209503:" + value["name"]
	k:=0
	for _, t := range alltopics.Topics {
		if deltemp == *t.TopicArn {
			fmt.Println("Topic already exists")
			fmt.Fprintln(w,"Topic exists")
			k=1
			break
		}
		}

	if k!=1{
		result, err := svc.CreateTopic(&sns.CreateTopicInput{
		Name: aws.String(value["name"]),
		})
		fmt.Println("Topic created ",*result.TopicArn)
		fmt.Fprintln(w, *result.TopicArn)
		if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		}
	}
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	}

/*
Documentation:

Constants::These are the credentials provided by AWS for a specific user account
	{
	AKID string = Access Key Identifier
	SECRET_KEY string = Secret Key
	}

func Home():: displays the homepage
func GetTopics():: displays all the existing topics
func CreateTopic():: creates a topic **Authentication is required
func GetSubByTopic():: shows all the subscriptions for the specific topic
func DeleteTopicByName():: deletes the specific topic
func CreateSub():: creates a subscription associated with a specific topic
func SendMsg():: sends the message to the specific topic

'''
Parameters:
1. http.ResponseWriter :: A ResponseWriter interface is used by an HTTP handler to construct an HTTP response.
2. *http.Request :: A Request represents an HTTP request received by a server or to be sent by a client.
'''

func NewSession():
	NewSession returns a new Session created from SDK defaults,
	config files, environment, and user provided config files.
	Once the Session is created it can be mutated to modify the Config or Handlers.

func WriteHeader():
	WriteHeader sends an HTTP response header with the provided status code

func Vars():
	Vars returns the route variables for the current request
*/