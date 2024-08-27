package handler

import (
	"github.com/gin-gonic/gin"
)

const html = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>ZKP Range Proofer</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
        }
        h1 {
            text-align: center;
        }
        .info {
            text-align: center;
            color: #666;
            margin-bottom: 20px;
        }
        form {
            display: flex;
            flex-direction: column;
            gap: 10px;
        }
        label {
            font-weight: bold;
        }
        input[type="text"], input[type="number"], input[type="file"] {
            width: 100%;
            padding: 5px;
        }
        button {
            padding: 10px;
            background-color: #4CAF50;
            color: white;
            border: none;
            cursor: pointer;
        }
        button:hover {
            background-color: #45a049;
        }
        #result {
            margin-top: 20px;
            border: 1px solid #ddd;
            padding: 10px;
            display: none;
        }
        .true {
            color: green;
        }
        .false {
            color: red;
        }
    </style>
</head>
<body>
    <h1>ZKP Range Proofer</h1>
    <p class="info">All operations are performed locally on your PC, ensuring the integrity and security of your proofs.</p>
    <form id="zkpForm">
        <label for="actualValue">Actual Value:</label>
        <input type="number" id="actualValue" name="actualValue" required>

        <label for="proofFile">Proof File:</label>
        <input type="file" id="proofFile" name="proofFile" accept=".json" required>
        
        <button type="submit">Verify Proof</button>
    </form>

    <div id="result">
        <h2>Proof Result:</h2>
        <p><strong>MinValue:</strong> <span id="minValue"></span></p>
        <p><strong>MaxValue:</strong> <span id="maxValue"></span></p>
        <p><strong>Actual Value:</strong> <span id="actualValueResult"></span></p>
        <p><strong>Result verification:</strong> <span id="verification"></span></p>
    </div>

    <script>
        document.getElementById('zkpForm').addEventListener('submit', function(e) {
            e.preventDefault();
            const fileInput = document.getElementById('proofFile');
            const actualValue = document.getElementById('actualValue').value;
            const file = fileInput.files[0];
            if (file) {
                const reader = new FileReader();
                reader.onload = function(e) {
                    const fileContent = e.target.result;
                    try {
                        const proofResult = JSON.parse(fileContent);
                        const requestBody = {
                            actual_value: parseInt(actualValue),
                            proof: proofResult
                        };
                        
                        fetch('/proof', {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json',
                            },
                            body: JSON.stringify(requestBody),
                        })
                        .then(response => response.json())
                        .then(data => {
                            document.getElementById('result').style.display = 'block';
                            document.getElementById('actualValueResult').textContent = data.actual_value;
                            document.getElementById('minValue').textContent = data.min_value;
                            document.getElementById('maxValue').textContent = data.max_value;
                            const verificationElement = document.getElementById('verification');
                            verificationElement.textContent = data.valid;
                            verificationElement.className = data.valid ? 'true' : 'false';
                        })
                        .catch((error) => {
                            console.error('Error:', error);
                            alert('An error occurred while verifying the proof.');
                        });
                    } catch (error) {
                        console.error('Error parsing JSON:', error);
                        alert('Invalid JSON file. Please upload a valid proof result file.');
                    }
                };
                reader.readAsText(file);
            }
        });
    </script>
</body>
</html>

`

func UI(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	ctx.String(200, html)
}
