# Logo Setup Instructions

## IMPORTANT: Add Your Logo Image

The website is configured to use your logo at:
```
/static/images/logo.png
```

### Steps to Add Your Logo:

1. **Save the logo image** you provided (the FlowStateCloud.ai/FlowStateLogic.ai logo with the dark blue background, cloud, and hexagonal pattern) as `logo.png`

2. **Place it in the images directory**:
   ```bash
   cp /path/to/your/logo.png /home/user/flowstatelogicai-website/static/images/logo.png
   ```

3. **Rebuild the application** (if using embedded files):
   ```bash
   cd /home/user/flowstatelogicai-website
   go build -o flowstatelogic-server main.go
   ```

### Logo Specifications:

- **Height**: 60px (configured in CSS)
- **Format**: PNG with transparent background recommended
- **Dimensions**: Your logo appears to be approximately 1200x600px - this will be scaled automatically
- **Background**: The logo's dark blue background matches the site theme perfectly

### Current Setup:

The website now has:
- **Hexagonal background pattern** throughout the site matching your logo
- **Animated floating hexagons** in the hero section
- **Dark blue color scheme** matching your logo exactly:
  - Background: rgb(20, 50, 70)
  - Primary: rgb(110, 180, 230)
  - Accents: rgb(160, 200, 240)
  - Cards: rgb(25, 60, 85)

### Testing:

After adding your logo, run the server and check:
```bash
go run main.go
```

Visit http://localhost:8080 - your logo should appear in the top-left of the navigation bar.

### Troubleshooting:

If the logo doesn't appear:
1. Check the file path is exactly: `/home/user/flowstatelogicai-website/static/images/logo.png`
2. Verify the file is a valid PNG image
3. Rebuild the Go application to embed the new file
4. Check browser console for any 404 errors
