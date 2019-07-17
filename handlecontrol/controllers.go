package handlecontrol

import (
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


var Home = func (w http.ResponseWriter, r *http.Request)(){
	if r.URL.Path != "/"{
		error(w,r, http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Welcome to Homepage")
}

var GetTopics = func(w http.ResponseWriter, r *http.Request)(){

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

var error = func(w http.ResponseWriter, r *http.Request, err int) {
	//WriteHeader sends an HTTP response header with the provided status code
	w.WriteHeader(err)
	if err == http.StatusNotFound {
		fmt.Fprint(w, "Error 404")
	}
}

var CreateTopic = func(w http.ResponseWriter, r *http.Request){
	value:= mux.Vars(r)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials(AKID, SECRET_KEY, ""),
	},)
	svc := sns.New(sess)
	result, err := svc.CreateTopic(&sns.CreateTopicInput{
		Name: aws.String(value["name"]),
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Topic created ",*result.TopicArn)
	fmt.Fprintln(w, *result.TopicArn)
}

var GetSubByTopic = func(w http.ResponseWriter, r *http.Request){
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

var DeleteTopicByName= func(w http.ResponseWriter, r *http.Request){
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

var CreateSub= func(w http.ResponseWriter, r *http.Request){

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

var SendMsg= func(w http.ResponseWriter, r *http.Request){

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
