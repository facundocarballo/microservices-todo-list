package com.facundocarballo.microservices.app.controllers;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import com.facundocarballo.microservices.app.AppConfig;
import com.facundocarballo.microservices.domain.User;
import com.facundocarballo.microservices.app.services.AuthService;

@RestController
public class AuthController {
    
    @PostMapping("/auth")
    public String handlePostAuth(@RequestBody String requestBody) {
        AppConfig appConfig = new AppConfig();
        AuthService service = new AuthService(appConfig.webClientBuilder(), "http://localhost:3001");
        User user = User.bodyToUser(requestBody);
        User realUser =service.getRealUser(user.getId());
        if (user.compareTo(realUser) == 0) {
            // Generate a JWT.
            return "Correct user credentials.";
        }
        return "Incorrect user credentials.";
    }
}
