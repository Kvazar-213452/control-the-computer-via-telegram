<?php

header('Content-Type: application/json');

$filename = '../data/foto.json';

if (!file_exists($filename)) {
    http_response_code(500);
    echo json_encode(['error' => 'Error reading the file']);
    exit;
}

$data = file_get_contents($filename);
if ($data === false) {
    http_response_code(500);
    echo json_encode(['error' => 'Error reading the file']);
    exit;
}

echo $data;

?>
