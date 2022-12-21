import path from 'path';

import express, { application } from 'express';
import bodyParser from 'body-parser';
import cors from 'cors';
// meedalware
app.use(
  cors(),
  bodyParser.urlencoded({ extended: true }),
  bodyParser.json()
)