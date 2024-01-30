package com.facundocarballo.microservices.app;

import com.facundocarballo.microservices.domain.EnviromentVariables;
import com.facundocarballo.microservices.domain.JWT;
import com.facundocarballo.microservices.domain.User;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class App {
    public static void main(String[] args) {
        SpringApplication.run(App.class, args);
        // EnviromentVariables env = new EnviromentVariables();
        // User user = new User(1, "facundocarballo", "facundo00");
        // JWT jwt = new JWT(user, env);

        // System.out.println("JWT Token: " + jwt.getToken());
    }
}
