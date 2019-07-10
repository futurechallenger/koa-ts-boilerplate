#!/usr/bin/env node
import Request from 'request';
import fs from 'fs';

const uri = '';
const basePath = '';
const FILE_NAME = 'go-cli';

function error(message: string): void {
  console.error(message);
}

const download = (uri: string, callback: () => void) => {
  Request.head(uri, (err, res, body) => {
    const r = Request(uri).pipe(fs.createWriteStream(basePath + FILE_NAME));
    r.on('close', callback);
    r.on('error', error);
  });
};

download(uri, () => {
  console.log('Download succeeded');
});
