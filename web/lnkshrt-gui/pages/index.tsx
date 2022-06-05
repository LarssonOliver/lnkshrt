import {
  Button,
  Container,
  Grid,
  Link,
  TextField,
  Typography,
  IconButton,
} from "@mui/material";
import ArrowRight from "@mui/icons-material/ArrowRightAlt";
import CopyButton from "@mui/icons-material/ContentCopy";
import GithubButton from "@mui/icons-material/GitHub";
import type { NextPage } from "next";
import Head from "next/head";
import { ChangeEventHandler, useEffect, useState } from "react";
import LinkModel from "../models/link";
import styles from "../styles/Home.module.css";
import axios from "axios";

const Home: NextPage = () => {
  const [url, setUrl] = useState("");
  const [shortUrl, setShortUrl] = useState<LinkModel | undefined>();
  const [apiUrl, setApiUrl] = useState("");

  useEffect(() => {
    axios.get("/path").then((res) => setApiUrl(res.data.path));
  }, []);

  const onChange: ChangeEventHandler<HTMLInputElement> = (e) => {
    const val = e.target.value;
    setUrl(val);
  };

  const onShorten = async () => {
    const response = await axios.post("/api", { url });
    setShortUrl(response.data);
  };

  const onCopy = async () => {
    if ("clipboard" in navigator && shortUrl)
      return await navigator.clipboard.writeText(`${apiUrl}/${shortUrl.id}`);
  };

  return (
    <div className={styles.container}>
      <Head>
        <title>lnkshrt</title>
        <meta name="description" content="an open source url shortener" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <Typography variant="h1" component="div">
          lnkshrt
        </Typography>

        <Typography variant="h5" component="div" marginBottom={"2em"}>
          an open source url shortener
        </Typography>

        <Container maxWidth="md">
          <Grid container spacing={2}>
            <Grid item xs={12} sm={9} md={10}>
              <TextField
                label="URL"
                fullWidth
                value={url}
                onChange={onChange}
                onKeyPress={(e) => {
                  if (e.key === "Enter") {
                    onShorten();
                  }
                }}
              />
            </Grid>
            <Grid item xs={12} sm={3} md={2}>
              <Button
                variant="outlined"
                className={styles.shortenButton}
                fullWidth
                onClick={onShorten}
                disabled={url === ""}
              >
                Shorten
              </Button>
            </Grid>
          </Grid>

          {shortUrl !== undefined && (
            <Container maxWidth="md" className={styles.resultContainer}>
              <Typography
                variant="h6"
                component="div"
                style={{ display: "flow-root", marginBottom: ".5rem" }}
              >
                <div style={{ float: "left" }}>
                  <Link
                    href={`${apiUrl}/${shortUrl.id}`}
                    underline="none"
                    className={styles.link}
                  >
                    {`${apiUrl}/${shortUrl.id}`}
                  </Link>
                </div>
                <Button
                  variant="outlined"
                  style={{ float: "right" }}
                  onClick={onCopy}
                >
                  Copy link
                  <CopyButton className={styles.copyButton} />
                </Button>
              </Typography>

              <Typography
                variant="body2"
                component="div"
                className={styles.longLink}
              >
                <ArrowRight color="disabled" className={styles.linkArrow} />
                {shortUrl.url}
              </Typography>
            </Container>
          )}
        </Container>

        <IconButton
          className={styles.github}
          href="https://github.com/larssonoliver/lnkshrt"
          target={"_blank"}
        >
          <GithubButton />
        </IconButton>
      </main>
    </div>
  );
};

export default Home;
