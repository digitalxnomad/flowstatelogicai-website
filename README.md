# FlowStateLogic.ai - Go Web Application

A modern, responsive web application built with Go that replicates the FlowStateLogic.ai website.

## Features

- **Pure Go HTTP Server**: Lightweight and efficient web server using Go's standard library
- **Embedded Static Files**: Static files are embedded into the binary using Go's embed directive
- **Responsive Design**: Mobile-first responsive design that works on all devices
- **Modern UI**: Clean, professional interface with smooth animations and interactions
- **Fast Performance**: Optimized CSS and JavaScript for quick page loads

## Tech Stack

- **Backend**: Go (Golang)
- **Frontend**: HTML5, CSS3, JavaScript
- **Font**: Lato (Google Fonts)
- **Design**: Custom CSS with modern layout techniques (Flexbox, Grid)

## Project Structure

```
flowstatelogicai-website/
├── main.go                 # Main Go server file
├── go.mod                  # Go module file
├── templates/
│   └── index.html         # Main HTML template
└── static/
    ├── css/
    │   └── styles.css     # Main stylesheet
    └── js/
        └── main.js        # JavaScript for interactivity
```

## Installation & Running

### Prerequisites

- Go 1.16 or higher (for embed support)

### Steps

1. **Clone or navigate to the repository**:
   ```bash
   cd flowstatelogicai-website
   ```

2. **Run the application**:
   ```bash
   go run main.go
   ```

   Or build and run the binary:
   ```bash
   go build -o flowstatelogic-server main.go
   ./flowstatelogic-server
   ```

3. **Access the website**:
   Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

## Configuration

The server runs on port 8080 by default. To change the port, modify the `port` variable in `main.go`:

```go
port := ":8080"  // Change to your desired port
```

## Features Implemented

### Sections

1. **Navigation Bar**
   - Responsive navigation with mobile menu
   - Smooth scroll to sections
   - Sticky header on scroll

2. **Hero Section**
   - Eye-catching headline with call-to-action
   - Gradient background
   - Responsive buttons

3. **Services Section**
   - Grid layout of service cards
   - Hover animations
   - SVG icons
   - Six key services:
     - Machine Learning
     - Natural Language Processing
     - Data Analytics
     - AI Consulting
     - Custom AI Solutions
     - AI Integration

4. **About Section**
   - Company information
   - Key statistics display
   - Visual graphics

5. **Contact Section**
   - Contact form with validation
   - Contact information display
   - Responsive layout

6. **Footer**
   - Quick links
   - Service links
   - Social media links
   - Copyright information

### Interactive Features

- Smooth scrolling navigation
- Mobile menu toggle
- Form submission handling
- Scroll-based animations
- Parallax effects
- Intersection Observer animations

## Design Specifications

### Color Palette

- **Primary**: RGB(242, 27, 63) - Vibrant pink/red
- **Secondary**: RGB(89, 89, 89) - Neutral gray
- **Info**: RGB(0, 151, 215) - Bright blue
- **Dark Background**: RGB(17, 21, 23)
- **Light Background**: RGB(255, 255, 255)

### Typography

- **Font Family**: Lato, Helvetica, Arial, sans-serif
- **Weights**: 300 (Light), 400 (Regular), 700 (Bold), 900 (Black)

### Responsive Breakpoints

- Mobile: < 640px
- Tablet: 640px - 1024px
- Desktop: > 1024px

## Deployment

### Build for Production

```bash
go build -ldflags="-s -w" -o flowstatelogic-server main.go
```

The `-ldflags="-s -w"` flags reduce binary size by stripping debug information.

### Docker Deployment (Optional)

Create a `Dockerfile`:

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w" -o server main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
```

Build and run:
```bash
docker build -t flowstatelogic-app .
docker run -p 8080:8080 flowstatelogic-app
```

## Development

### Adding New Pages

1. Create a new HTML file in the `templates/` directory
2. Add a new route handler in `main.go`:

```go
http.HandleFunc("/new-page", func(w http.ResponseWriter, r *http.Request) {
    data, _ := templateFiles.ReadFile("templates/new-page.html")
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.Write(data)
})
```

### Modifying Styles

Edit `static/css/styles.css` to customize the appearance. The CSS uses CSS custom properties (variables) for easy theme customization.

### Adding JavaScript Functionality

Edit `static/js/main.js` to add new interactive features.

## Performance Optimizations

- Embedded static files eliminate file system reads
- Minified CSS and optimized selectors
- Efficient JavaScript with event delegation
- Lazy loading animations with Intersection Observer
- CSS Grid and Flexbox for optimal layout performance

## Browser Support

- Chrome (latest)
- Firefox (latest)
- Safari (latest)
- Edge (latest)
- Mobile browsers (iOS Safari, Chrome Mobile)

## License

Copyright © 2025 FlowStateLogic.ai. All rights reserved.

## Contact

For questions or support, visit the contact section on the website or email: contact@flowstatelogic.ai