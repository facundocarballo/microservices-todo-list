package com.facundocarballo.microservices.domain;

public class JWT {

    private String token;

    public JWT(User user, EnviromentVariables env) {
        this.token = "";
    }

    public String getToken() {
        return this.token;
    }
}
