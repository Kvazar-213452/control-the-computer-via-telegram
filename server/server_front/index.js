const express = require('express');
const path = require('path');
const multer = require('multer');
const fs = require('fs');

const app = express();
const PORT = process.env.PORT || 3000;

const storage = multer.diskStorage({
  destination: (req, file, cb) => {
    cb(null, path.join(__dirname, 'files'));
  },
  filename: (req, file, cb) => {
    const uniqueSuffix = Date.now();
    const extension = path.extname(file.originalname);
    const originalName = path.basename(file.originalname, extension);
    cb(null, `${originalName}-${uniqueSuffix}${extension}`);
  },
});

app.post('/data', (req, res) => {
  const { name, age } = req.body;

  if (!name || !age) {
    return res.status(400).json({ error: 'Name and age are required.' });
  }

  const filePath = path.join(__dirname, 'main.json');
  res.sendFile(filePath);
});

const upload = multer({ storage });

app.set('view engine', 'ejs');
app.set('views', path.join(__dirname, 'views'));
app.use(express.static(path.join(__dirname, 'public')));

app.get('/', (req, res) => {
  res.render('index');
});

app.post('/upload', upload.single('file'), (req, res) => {
  const text = req.body.text;
  let fff = text;
  const file = req.file;

  if (!file || !text) {
    return res.status(400).send('Файл або текст не надано.');
  }

  const jsonFilePath = path.join(__dirname, 'main.json');

  let jsonData = [];
  if (fs.existsSync(jsonFilePath)) {
    const fileContent = fs.readFileSync(jsonFilePath, 'utf-8');
    jsonData = JSON.parse(fileContent);
  }

  jsonData.push([fff, file.originalname]);
  
  fs.writeFileSync(jsonFilePath, JSON.stringify(jsonData, null, 2), 'utf-8');

  res.send(`Файл "${file.originalname}" і текст "${text}" успішно додано до першого об'єкта JSON.`);
});

app.use((req, res) => {
  res.status(404).render('404', { title: '404 - Сторінку не знайдено' });
});

app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
