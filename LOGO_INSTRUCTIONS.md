# Logo Setup Instructions

## Important: Add Your Logo

The website is configured to use a logo image at:
```
/static/images/logo.png
```

**You provided a logo image in our conversation. Please save it to the path above.**

### Steps to Add the Logo:

1. Save the logo image you provided (FlowStateCloud.ai logo with hexagonal cloud pattern) as `logo.png`
2. Place it in: `/home/user/flowstatelogicai-website/static/images/logo.png`

### Logo Specifications:
- The logo will be displayed at 50px height in the navigation bar
- The logo should have a transparent background for best results
- The current design uses the dark blue theme from your logo (rgb(20, 50, 70))

### Alternative: Use Text-Only Logo

If you prefer not to use an image logo, you can remove the `<img>` tag from the navigation in `/templates/index.html` line 18, and the text logo will be used instead.

### Current Color Scheme

The website now uses the color scheme from your logo:
- **Background**: Dark blue rgb(20, 50, 70)
- **Primary**: Light blue rgb(110, 180, 230)
- **Accents**: Secondary blue rgb(160, 200, 240)
- **Cards**: Medium blue rgb(25, 60, 85)
