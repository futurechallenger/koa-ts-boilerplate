#!/usr/bin/env node
import Request from 'request';
import process from 'process';
import fs from 'fs';
import path from 'path';
import decompress from 'decompress';
import { FILE } from 'dns';

const uri =
  'https://github.com/futurechallenger/koa-ts-boilerplate/releases/download/v1.0.3/koa-ts-boilerplate_1.0.3_Darwin_x86_64.tar.gz';
let basePath = process.cwd();
const FILE_NAME = 'go-cli';
const DEST_FILE_PREFIX = '.raw';

function error(message: string): void {
  console.error(message);
}

const download = (uri: string, callback: (destPath: string) => () => void) => {
  const destPath = path.join(basePath, `${FILE_NAME}${DEST_FILE_PREFIX}`);
  Request.head(uri, (err, res, body) => {
    const r = Request(uri).pipe(fs.createWriteStream(destPath));
    r.on('close', callback(destPath));
    r.on('error', error);
  });
};

download(uri, (destPath: string) => () => {
  fs.access(path.join(basePath, FILE_NAME), fs.constants.F_OK, err => {
    if (!err) {
      fs.unlink(path.join(destPath, FILE_NAME), err => {
        console.error('Existing files does not removed');
      });
    }
    decompress(destPath, basePath).then(files =>
      console.log('decompressed file: ', files),
    );
  });
});
