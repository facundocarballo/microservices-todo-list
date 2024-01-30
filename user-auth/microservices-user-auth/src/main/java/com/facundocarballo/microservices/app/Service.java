package com.facundocarballo.microservices.app;

import org.springframework.web.reactive.function.client.WebClient;

import com.facundocarballo.microservices.domain.User;

import reactor.core.publisher.Mono;

public class Service {
 
    private final WebClient webClient;

    public Service(WebClient.Builder webClientBuilder) {
        this.webClient = webClientBuilder.baseUrl("http://localhost:3001").build();
    }

    public User obtenerDatosDeOtraAPI() {
        // Realizar una solicitud GET
        Mono<User> respuestaMono = webClient.get()
                .uri("/user/1")
                .retrieve()
                .bodyToMono(User.class);

        // Bloquear hasta que se complete la solicitud y obtener el resultado
        User respuesta = respuestaMono.block();

        return respuesta;
    }
}
