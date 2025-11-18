<?php
$db = new mysqli("localhost", "root", "rootpassword", "testdb");

$username = $_GET["user"];
$password = $_GET["pass"];

$query = "SELECT * FROM users WHERE username = '$username' AND password = '$password'";
$result = $db->query($query);

if ($result && $result->num_rows > 0) {
    echo "Login successful!";
} else {
    echo "Invalid login.";
}
?>
