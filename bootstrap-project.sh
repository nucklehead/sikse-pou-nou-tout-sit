#!/bin/bash

set -x

END_OF_DEV_SECTION="log.request.output = stderr"
END_OF_PROD_SECTION="log.request.output = off"

DATABASE_CONFIG="\n\nmongo.database = heroku_blh7zd3p\nmongo.path = ds149049.mlab.com:49049\nmongo.maxPool = 20"

REVEL_MGO_FUNCTION="

func InitDB() {
	mongodb.MaxPool = revel.Config.IntDefault(\"mongo.maxPool\", 0)
	mongodb.PATH,_ = revel.Config.String(\"mongo.path\")
	mongodb.DBNAME, _ = revel.Config.String(\"mongo.database\")
	mongodb.CheckAndInitServiceConnection()
	log.Print(mongodb.DBNAME)
}
"

GITHUB="github.com"

BAD_IMPORT="$2"

BAD_RENDER_METHOD="RenderJson"
GOOD_RENDER_METHOD="RenderJSON"

display_usage() {
	echo "Incorrect inputs."
	echo -e "\nUsage:\n./bootstrap-project.sh <github_account> <project_name> \n"
}

if [  $# -ne 2 ]
    then
		display_usage
		exit 1
fi

if [[ ( $# == "--help") ||  $# == "-h" ]]
	then
		display_usage
		exit 0
fi


# Dependencies ####################################################################

#go get github.com/revel/revel
#go get github.com/revel/cmd/revel
#go get github.com/kyawmyintthein/revel_mgo


# Setup ####################################################################

#revel new github.com/$1/$2
#revel_mgo mgo:setup


# Configuration ####################################################################

sed -i "s/\($END_OF_DEV_SECTION\)/\1$DATABASE_CONFIG/" $GOPATH/src/github.com/$1/$2/conf/app.conf

sed -i "s/\($END_OF_PROD_SECTION\)/\1$DATABASE_CONFIG/" $GOPATH/src/github.com/$1/$2/conf/app.conf

cat >> $GOPATH/src/github.com/$1/$2/app/init.go <<EOF
$REVEL_MGO_FUNCTION
EOF

# Models Code generation ####################################################################

revel_mgo generate model Account -fields=username:string,password:string,email:string,phone:string
revel_mgo generate model Comment -fields=user:string,content:string
revel_mgo generate model Event -fields=title:string,date:datetime,description:string,speaker:string,location:string
revel_mgo generate model Option -fields=name:string,description:string
revel_mgo generate model Presenter -fields=firstName:string,lastName:string,twitter:string,about:string,location:string,email:string,phone:string
revel_mgo generate model Sponsor -fields=name:string,description:string
revel_mgo generate model Video -fields=link:string,title:string,description:string

find  $GOPATH/src/github.com/$1/$2/app/models -type f -exec sed -i "s/\($BAD_IMPORT\)/$GITHUB\/$1\/\1/g" {} \;


# Controllers Code generation ####################################################################

revel_mgo generate controller Account
revel_mgo generate controller Comment
revel_mgo generate controller Event
revel_mgo generate controller Option
revel_mgo generate controller Presenter
revel_mgo generate controller Sponsor
revel_mgo generate controller Video

find  $GOPATH/src/github.com/$1/$2/app/controllers -type f -exec sed -i "s/\($BAD_IMPORT\)/$GITHUB\/$1\/\1/g" {} \;
find  $GOPATH/src/github.com/$1/$2/app/controllers -type f -exec sed -i "s/$BAD_RENDER_METHOD/$GOOD_RENDER_METHOD/g" {} \;

# Routes configuration ####################################################################

MODELS="Account
Comment
Event
Option
Presenter
Sponsor
Video"

for model in $MODELS; do
    cat >> $GOPATH/src/github.com/$1/$2/conf/routes <<EOF

# $model
POST    /${model,,}/                             ${model}Controller.Create
GET     /${model,,}/                             ${model}Controller.Index
GET     /${model,,}/:id                          ${model}Controller.Show
PUT     /${model,,}/:id                          ${model}Controller.Update
DELETE  /${model,,}/:id                          ${model}Controller.Delete
GET     /show/${model,,}/                        ${model}Controller.ShowList
EOF
    done;