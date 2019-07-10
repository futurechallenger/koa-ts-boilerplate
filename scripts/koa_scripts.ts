#!/usr/bin/env node
import Request from 'request';
import process from 'process';
import fs from 'fs';
import path from 'path';

const uri =
  'https://github.com/futurechallenger/koa-ts-boilerplate/releases/download/v1.0.3/koa-ts-boilerplate_1.0.3_Darwin_i386.tar.gz';
let basePath = process.cwd();
const FILE_NAME = 'go-cli';

function error(message: string): void {
  console.error(message);
}

const download = (uri: string, callback: (destPath: string) => () => void) => {
  const destPath = path.join(basePath, FILE_NAME);
  console.log('===>dest path: ', destPath);
  Request.head(uri, (err, res, body) => {
    const r = Request(uri).pipe(fs.createWriteStream(destPath));
    r.on('close', callback(destPath));
    r.on('error', error);
  });
};

download(uri, (destPath: string) => () => {
  console.log('Download succeeded');
});
