# Configure Enviroment Variables
$env:JWT_SECRET_KEY = "FacundoCarballo"

# Execute our Java Program
mvn clean install
java -jar target/microservices-user-auth-1.0.jar