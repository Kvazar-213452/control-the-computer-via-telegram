<?php

header('Content-Type: application/json');

$uploadDir = '../public/file/';
$filename = '../data/music.json';
$serverUrl = 'http://localhost/';

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    if (!isset($_FILES['file']) || !isset($_POST['text'])) {
        http_response_code(400);
        echo json_encode(['error' => 'Missing file or text parameter.']);
        exit;
    }

    $text = $_POST['text'];
    $file = $_FILES['file'];
    $filePath = $uploadDir . basename($file['name']);
    
    if (!move_uploaded_file($file['tmp_name'], $filePath)) {
        http_response_code(500);
        echo json_encode(['error' => 'Failed to save file.']);
        exit;
    }

    $jsonData = [];
    if (file_exists($filename)) {
        $fileContent = file_get_contents($filename);
        $jsonData = json_decode($fileContent, true) ?? [];
    }

    $jsonData[$text] = $serverUrl . "file/" . basename($file['name']);
    file_put_contents($filename, json_encode($jsonData, JSON_PRETTY_PRINT));
    
    echo json_encode(['message' => 'ok']);
    exit;
}

http_response_code(405);
echo json_encode(['error' => 'Method not allowed.']);
exit;

?>
