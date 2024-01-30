package com.facundocarballo.microservices;

import com.facundocarballo.microservices.domain.EnviromentVariables;
import com.facundocarballo.microservices.domain.JWT;
import com.facundocarballo.microservices.domain.User;

public class App {
    public static void main(String[] args) {
        EnviromentVariables env = new EnviromentVariables();
        User user = new User(1, "facundocarballo", "facundo00");
        JWT jwt = new JWT(user, env);

        System.out.println("JWT Token: " + jwt.getToken());
    }
}
