const express = require('express');
const path = require('path');
const multer = require('multer');
const fs = require('fs');

const app = express();
const PORT = process.env.PORT || 3000;

const Server_url = "http://localhost:3000/";

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

app.use(express.json());
app.set('view engine', 'ejs');
app.set('views', path.join(__dirname, 'views'));
app.use(express.static(path.join(__dirname, 'public')));

app.get('/', (req, res) => {
  res.render('index');
});

app.get('/music', (req, res) => {
  fs.readFile('data/music.json', (err, data) => {
    if (err) {
      res.status(500).send('Error reading the file');
    } else {
      res.json(JSON.parse(data));
    }
  });
});

app.get('/foto', (req, res) => {
  fs.readFile('data/foto.json', (err, data) => {
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

  const jsonFilePath = path.join(__dirname, 'data/music.json');

  if (fs.existsSync(jsonFilePath)) {
    const fileContent = fs.readFileSync(jsonFilePath, 'utf-8');
    jsonData = JSON.parse(fileContent);
  }

  jsonData[text] = Server_url + "file/" + file.filename;

  fs.writeFileSync(jsonFilePath, JSON.stringify(jsonData, null, 2), 'utf-8');

  res.send("ok");
});

app.post('/upload_foto', upload.single('file'), (req, res) => {
  const text = req.body.text;
  const file = req.file;

  if (!file || !text) {
    return res.status(400).send('error.');
  }

  const jsonFilePath = path.join(__dirname, 'data/foto.json');

  if (fs.existsSync(jsonFilePath)) {
    const fileContent = fs.readFileSync(jsonFilePath, 'utf-8');
    jsonData = JSON.parse(fileContent);
  }

  jsonData[text] = Server_url + "file/" + file.filename;

  fs.writeFileSync(jsonFilePath, JSON.stringify(jsonData, null, 2), 'utf-8');

  res.send("ok");
});

app.post('/del_music', (req, res) => {
  const jsonFilePath = path.join(__dirname, 'data', 'music.json');
  const text_ = req.body.text;

  fs.readFile(jsonFilePath, 'utf-8', (err, data) => {
    if (err) {
      return res.status(500).json({ error: 'error read' });
    }

    let jsonData;
    try {
      jsonData = JSON.parse(data);
    } catch (parseError) {
      return res.status(500).json({ error: 'error JSON' });
    }

    if (jsonData[text_]) {
      const fileName = jsonData[text_].split('/').pop();
      const filePath = path.join(__dirname, 'public', 'file', fileName);

      fs.unlink(filePath, (unlinkError) => {
        if (unlinkError) {
          return res.status(500).json({ error: 'error deleting file' });
        }

        delete jsonData[text_];
        fs.writeFile(jsonFilePath, JSON.stringify(jsonData, null, 2), 'utf-8', (writeError) => {
          if (writeError) {
            return res.status(500).json({ error: 'error writing JSON' });
          }

          return res.json({ message: 'good' });
        });
      });
    } else {
      return res.status(404).json({ error: 'error not found' });
    }
  });
});

app.post('/del_foto', (req, res) => {
  const jsonFilePath = path.join(__dirname, 'data', 'foto.json');
  const text_ = req.body.text;

  fs.readFile(jsonFilePath, 'utf-8', (err, data) => {
    if (err) {
      return res.status(500).json({ error: 'error read' });
    }

    let jsonData;
    try {
      jsonData = JSON.parse(data);
    } catch (parseError) {
      return res.status(500).json({ error: 'error JSON' });
    }

    if (jsonData[text_]) {
      const fileName = jsonData[text_].split('/').pop();
      const filePath = path.join(__dirname, 'public', 'file', fileName);

      fs.unlink(filePath, (unlinkError) => {
        if (unlinkError) {
          return res.status(500).json({ error: 'error deleting file' });
        }

        delete jsonData[text_];
        fs.writeFile(jsonFilePath, JSON.stringify(jsonData, null, 2), 'utf-8', (writeError) => {
          if (writeError) {
            return res.status(500).json({ error: 'error writing JSON' });
          }

          return res.json({ message: 'good' });
        });
      });
    } else {
      return res.status(404).json({ error: 'error not found' });
    }
  });
});

app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
