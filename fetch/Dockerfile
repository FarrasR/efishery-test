FROM openjdk:17-alpine

COPY product.db product.db
COPY build/libs/fetch-api-0.0.1-SNAPSHOT.jar /app/application.jar

EXPOSE 8080

CMD ["java", "-jar", "/app/application.jar"]