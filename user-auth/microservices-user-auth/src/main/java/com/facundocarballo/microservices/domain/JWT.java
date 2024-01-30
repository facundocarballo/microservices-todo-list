package com.facundocarballo.microservices.domain;

import io.jsonwebtoken.Jwts;

public class JWT {

    private String token;

    public JWT(User user, EnviromentVariables env) {
        this.token = "";
    }

    public String getToken() {
        return this.token;
    }
}
