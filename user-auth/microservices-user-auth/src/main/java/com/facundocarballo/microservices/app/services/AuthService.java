package com.facundocarballo.microservices.app.services;

import org.springframework.web.reactive.function.client.WebClient;
import reactor.core.publisher.Mono;
import com.facundocarballo.microservices.domain.User;

public class AuthService {
    private final WebClient webClient;

    public AuthService(WebClient.Builder webClientBuilder, String url) {
        this.webClient = webClientBuilder.baseUrl(url).build();
    }

    // TODO: Implementar una nueva API desde el microservicio de Golang para que este microservicio pueda acceder a la informacion del usuario completa.
    // email, username, password.
    public User getRealUser(int userId) {
        Mono<User> res = this.webClient.get()
                            .uri("/user/" + userId)
                            .retrieve()
                            .bodyToMono(User.class);
        return res.block();
    }
}
