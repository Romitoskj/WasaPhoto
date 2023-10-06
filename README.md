# Wasa Photo

>Keep in touch with your friends by sharing photos of special moments, thanks to WASAPhoto! You can upload your photos directly from your PC, and they will be visible to everyone following you.

This repository contains the homework project for [Web and Software Architecture course](http://gamificationlab.uniroma1.it/en/wasa/) that I've taken at Sapienza University of Rome.

## Project structure

The project is structured in 4 parts:
- OpenApi specification
- Go backend
- Vue.js frontend
- Docker deployment

## Project template

This project is based on [Fantastic coffee (decaffeinated)](https://github.com/sapienzaapps/fantastic-coffee-decaffeinated) template. is a simplified version for the WASA course, not suitable for a production environment.
The full version can be found in the "Fantastic Coffee" repository.

## Build container images

### Backend
```
$ docker build -t wasa-photo-backend:latest -f Dockerfile.backend .
```

### Frontend
```
$ docker build -t wasa-photo-frontend:latest -f Dockerfile.frontend .
```

## Run container images

### Backend
```
$ docker run -it --rm -p 3000:3000 wasa-photo-backend:latest
```

### Frontend
```
$ docker run -it --rm -p 8080:80 wasa-photo-frontend:latest
```

## License

See [LICENSE](LICENSE).
