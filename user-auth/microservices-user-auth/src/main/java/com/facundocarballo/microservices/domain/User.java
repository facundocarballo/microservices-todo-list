package com.facundocarballo.microservices.domain;

import java.util.HashMap;
import java.util.Map;

public class User {
    protected int id;
    protected String username;
    protected String email;
    protected String password;

    public User(int id, String username, String password) {
        this.id = id;
        this.username = username;
        this.password = password;
    }

    public User(String email, String password, int id) {
        this.id = id;
        this.email = email;
        this.password = password;
    }

    public Map<String, Object> getJSON() {
        Map<String, Object> jsonPayload = new HashMap<>();
        
        jsonPayload.put("username", this.username);
        jsonPayload.put("email", this.email);
        jsonPayload.put("password", this.password);

        return jsonPayload;
    }
}
