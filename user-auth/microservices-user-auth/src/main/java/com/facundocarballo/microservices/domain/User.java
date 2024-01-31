package com.facundocarballo.microservices.domain;

import java.util.HashMap;
import java.util.Map;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.ObjectMapper;

public class User implements Comparable<User> {
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

    @JsonCreator
    public User(
            @JsonProperty("id") int id,
            @JsonProperty("username") String username,
            @JsonProperty("email") String email,
            @JsonProperty("password") String password) {
        this.id = id;
        this.username = username;
        this.email = email;
        this.password = password;
    }

    public static User bodyToUser(String body) {
        ObjectMapper objManager = new ObjectMapper();
        try {
            return objManager.readValue(body, User.class);
        } catch (Exception e) {
            System.err.println("Error casting the body request to User class. " + e.getStackTrace());
        }
        return null;
    }

    public Map<String, Object> getJSON() {
        Map<String, Object> jsonPayload = new HashMap<>();

        jsonPayload.put("username", this.username);
        jsonPayload.put("email", this.email);
        jsonPayload.put("password", this.password);

        return jsonPayload;
    }

    public int getId() {
        return this.id;
    }

    public String getUsername() {
        return this.username;
    }

    public String getEmail() {
        return this.email;
    }

    public String getPassword() {
        return this.password;
    }

    @Override
    public int compareTo(User user) {
        if (!this.email.equals(user.getEmail()) ||
                !this.username.equals(user.getUsername()))
            return 1;

        if (!this.password.equals(user.getPassword()))
            return 1;

        return 0;
    }

    @Override
    public boolean equals(Object obj) {
        return super.equals(obj);
    }

    @Override
    public int hashCode() {
        return super.hashCode();
    }
}
