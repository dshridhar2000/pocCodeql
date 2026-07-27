package com.example.vulnerable;

import java.io.IOException;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.Statement;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class VulnerableServlet extends HttpServlet {
    protected void doGet(HttpServletRequest request, HttpServletResponse response) 
            throws ServletException, IOException {
        
        // CodeQL registers this parameter as untrusted source data ("Taint Source")
        String customerId = request.getParameter("id"); 

        try {
            Connection conn = DriverManager.getConnection("jdbc:mysql://localhost:3306/db", "user", "pass");
            Statement stmt = conn.createStatement();

            // ALERT: SQL Injection (java/sql-injection)
            // Tainted data is concatenated into a query string and passed straight into a database execution sink
            String query = "SELECT * FROM customers WHERE id = '" + customerId + "'";
            ResultSet rs = stmt.executeQuery(query);

            if (rs.next()) {
                response.getWriter().println("User found: " + rs.getString("name"));
            }
            conn.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
