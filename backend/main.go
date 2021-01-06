package main
 
import (
   "context"
   "log"
 
   "github.com/panupong/app/controllers"
   _ "github.com/panupong/app/docs"
   "github.com/panupong/app/ent"
   "github.com/gin-contrib/cors"
   "github.com/gin-gonic/gin"
   _ "github.com/mattn/go-sqlite3"
   swaggerFiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"

   
)

type Genders struct{
    Gender []Gender
}
type Gender struct{
    GenderType string
}

type Titles struct{
    Title []Title
}

type Title struct{
    TitleType string
}

type Jobs struct{
    Job []Job
}

type Job struct{
    JobType string
}



// @title SUT SA Example API Patient
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
    router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

 
   v1 := router.Group("/api/v1")
   controllers.NewGenderController(v1, client)
   controllers.NewTitleController(v1, client)
   controllers.NewJobController(v1, client)
   controllers.NewPatientController(v1, client)

   //set Genders Data
   genders := Genders{
       Gender: []Gender{
           Gender{"ชาย"},
           Gender{"หญิง"},
           Gender{"ไม่ระบุ"},
       },
   }
   for _, g := range genders.Gender{
    client.Gender.
     Create().
     SetGenderType(g.GenderType).Save(context.Background())
    }

   //set Title Data
   title := Titles{
       Title: []Title{
           Title{"เด็กชาย"},
           Title{"เด็กหญิง"},
           Title{"นาย"},
           Title{"นางสาว"},
           Title{"ไม่ระบุ"},
       },
   }
   for _, t := range title.Title{
    client.Title.
    Create().
    SetTitleType(t.TitleType).Save(context.Background()) 
    }


   //set Job Data
   job := Jobs{
       Job: []Job{
           Job{"เกษตรกร"},
           Job{"ผู้ใหญ่บ้าน"},
           Job{"ข้าราชการ"},
           Job{"การประมง‎"},
           Job{"พนักงานเอกชน"},
           Job{"ไม่ระบุ"},
       },
   }
   for _, j := range job.Job{
    client.Job.
    Create().
    SetJobName(j.JobType).Save(context.Background())
    }

 
   router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
   router.Run()
}