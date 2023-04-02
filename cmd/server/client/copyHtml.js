// Copies all html files from src to dist, used by npm run build
const fs = require('fs');
const path = require('path');

const srcDir = 'src';
const distDir = 'dist';

function copyHtmlFiles(src, dest) {
  const files = fs.readdirSync(src);

  files.forEach(file => {
    const srcFile = path.join(src, file);
    const destFile = path.join(dest, file);
    const stat = fs.lstatSync(srcFile);

    if (stat.isDirectory()) {
      if (!fs.existsSync(destFile)) {
        fs.mkdirSync(destFile);
      }
      copyHtmlFiles(srcFile, destFile);
    } else if (path.extname(srcFile) === '.html') {
      fs.copyFileSync(srcFile, destFile);
    }
  });
}

copyHtmlFiles(srcDir, distDir);