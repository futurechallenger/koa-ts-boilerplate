#!/usr/bin/env node
import Request from 'request';
import process from 'process';
import fs from 'fs';
import path from 'path';
import decompress from 'decompress';

const uri =
  'https://github.com/futurechallenger/koa-ts-boilerplate/releases/download/v1.0.3/koa-ts-boilerplate_1.0.3_Darwin_i386.tar.gz';
let basePath = process.cwd();
const FILE_NAME = 'go-cli.raw';
const DEST_FILE_PREFIX = '.raw';

function error(message: string): void {
  console.error(message);
}

const download = (uri: string, callback: (destPath: string) => () => void) => {
  const destPath = path.join(basePath, FILE_NAME);
  Request.head(uri, (err, res, body) => {
    const r = Request(uri).pipe(fs.createWriteStream(destPath));
    r.on('close', callback(destPath));
    r.on('error', error);
  });
};

download(uri, (destPath: string) => () => {
  console.log('Destination path: ', destPath);
  const destFilename = destPath.indexOf('.raw');
  decompress(destPath, 'go-cli').then(files => console.log('decompressed file: ', files));
});
