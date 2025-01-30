const express = require('express');
const path = require('path');
const multer = require('multer');
const fs = require('fs');

const app = express();
const PORT = process.env.PORT || 3000;


const storage = multer.diskStorage({
  destination: (req, file, cb) => {
    cb(null, path.join(__dirname, 'public/file'));
  },
  filename: (req, file, cb) => {
    const uniqueSuffix = Date.now();
    const extension = path.extname(file.originalname);
    const originalName = path.basename(file.originalname, extension);
    cb(null, `${originalName}-${uniqueSuffix}${extension}`);
  },
});
const upload = multer({ storage });

app.set('view engine', 'ejs');
app.set('views', path.join(__dirname, 'views'));
app.use(express.static(path.join(__dirname, 'public')));

app.get('/', (req, res) => {
  res.render('index');
});

app.get('/music', (req, res) => {
  fs.readFile('music.json', (err, data) => {
    if (err) {
      res.status(500).send('Error reading the file');
    } else {
      res.json(JSON.parse(data));
    }
  });
});

app.get('/foto', (req, res) => {
  fs.readFile('foto.json', (err, data) => {
    if (err) {
      res.status(500).send('Error reading the file');
    } else {
      res.json(JSON.parse(data));
    }
  });
});

app.post('/upload_music', upload.single('file'), (req, res) => {
  const text = req.body.text;
  const file = req.file;

  if (!file || !text) {
    return res.status(400).send('error.');
  }

  const jsonFilePath = path.join(__dirname, 'music.json');

  if (fs.existsSync(jsonFilePath)) {
    const fileContent = fs.readFileSync(jsonFilePath, 'utf-8');
    jsonData = JSON.parse(fileContent);
  }

  jsonData[text] = "http://localhost:3000/file/" + file.filename;

  fs.writeFileSync(jsonFilePath, JSON.stringify(jsonData, null, 2), 'utf-8');

  res.send("ok");
});

app.post('/upload_foto', upload.single('file'), (req, res) => {
  const text = req.body.text;
  const file = req.file;

  if (!file || !text) {
    return res.status(400).send('error.');
  }

  const jsonFilePath = path.join(__dirname, 'music.json');

  if (fs.existsSync(jsonFilePath)) {
    const fileContent = fs.readFileSync(jsonFilePath, 'utf-8');
    jsonData = JSON.parse(fileContent);
  }

  jsonData[text] = "http://localhost:3000/file/" + file.filename;

  fs.writeFileSync(jsonFilePath, JSON.stringify(jsonData, null, 2), 'utf-8');

  res.send("ok");
});

app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
