const express = require('express');
const app = express();

app.get('/execute', (req, res) => {
    // CodeQL registers this query string parameter as untrusted taint data
    const userCode = req.query.code;

    try {
        // ⚠️ INTENTIONAL VULNERABILITY: Code Injection
        // Passing user input straight into eval() executes any arbitrary JavaScript string
        const result = eval(userCode); 
        
        res.send(`Result: ${result}`);
    } catch (err) {
        res.status(500).send("Execution failed");
    }
});

app.listen(3000, () => console.log('Server running on port 3000'));
