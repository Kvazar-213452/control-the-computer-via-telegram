<?php

header('Content-Type: application/json');

$uploadDir = '../public/file/';
$filename = '../data/foto.json';
$serverUrl = 'http://localhost/';

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $input = json_decode(file_get_contents('php://input'), true);
    if (!isset($input['text'])) {
        http_response_code(400);
        echo json_encode(['error' => 'Missing text parameter.']);
        exit;
    }

    $text_ = $input['text'];

    if (!file_exists($filename)) {
        http_response_code(500);
        echo json_encode(['error' => 'Error reading file']);
        exit;
    }

    $fileContent = file_get_contents($filename);
    $jsonData = json_decode($fileContent, true);
    if (!is_array($jsonData)) {
        http_response_code(500);
        echo json_encode(['error' => 'Error parsing JSON']);
        exit;
    }

    if (isset($jsonData[$text_])) {
        $fileName = basename($jsonData[$text_]);
        $filePath = $uploadDir . $fileName;

        if (file_exists($filePath) && !unlink($filePath)) {
            http_response_code(500);
            echo json_encode(['error' => 'Error deleting file']);
            exit;
        }

        unset($jsonData[$text_]);

        if (file_put_contents($filename, json_encode($jsonData, JSON_PRETTY_PRINT)) === false) {
            http_response_code(500);
            echo json_encode(['error' => 'Error writing JSON']);
            exit;
        }

        echo json_encode(['message' => 'good']);
        exit;
    }

    http_response_code(404);
    echo json_encode(['error' => 'Error not found']);
    exit;
}

http_response_code(405);
echo json_encode(['error' => 'Method not allowed']);
exit;

?>
