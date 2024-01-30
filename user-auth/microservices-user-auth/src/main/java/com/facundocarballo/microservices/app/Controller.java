package com.facundocarballo.microservices.app;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import com.facundocarballo.microservices.domain.User;
import com.fasterxml.jackson.databind.ObjectMapper;

@RestController
public class Controller {

    @PostMapping("/auth")
    public String helloWorld(@RequestBody String requestBody) {
        ObjectMapper objectMapper = new ObjectMapper();
        AppConfig app = new AppConfig();
        Service service = new Service(app.webClientBuilder());
        try {
            User user = objectMapper.readValue(requestBody, User.class);
            User realUser = service.obtenerDatosDeOtraAPI();
            if (user.getEmail().equals(realUser.getEmail())) {
                return "Los emails son las mismos.";
            }
            return "Objeto Java creado a partir del JSON: " + user.toString();
        } catch (Exception e) {
            e.printStackTrace();
            return "Error al convertir el JSON a objeto Java";
        }
    }

}
