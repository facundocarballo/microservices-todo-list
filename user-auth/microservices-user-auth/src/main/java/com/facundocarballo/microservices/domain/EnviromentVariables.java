package com.facundocarballo.microservices.domain;

public class EnviromentVariables {
    private byte[] jwtSecretKey;

    public EnviromentVariables() {
        this.jwtSecretKey = System.getenv("JWT_SECRET_KEY").getBytes();
    }

    public byte[] GetJwtSecretKey() {
        return this.jwtSecretKey;
    } 
}
