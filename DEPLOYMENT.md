# 🦭 Seal ASCII Animation - Vercel Deployment Guide

This guide shows how to deploy your 187-frame braille seal animation to Vercel.

## Prerequisites

1. **Vercel CLI installed**:
   ```bash
   npm i -g vercel
   ```

2. **Vercel account**: Sign up at [vercel.com](https://vercel.com)

3. **Git repository**: Your code should be in a Git repository

## Step 1: Update Vercel Handler with 187 Frames

First, we need to update `api/index.go` to use your 187 braille frames instead of the simple 8 frames.

### Option A: Copy Frames Directly
Replace the `SealWiggleFrames` array in `api/index.go` with the frames from `animations/seal_wiggle.go`:

```bash
# Extract just the frames array from your main file
sed -n '/var SealWiggleFrames = \[\]string{/,/^}/p' animations/seal_wiggle.go > temp_frames.txt
```

Then manually copy the frames array into `api/index.go`, replacing lines 16-160.

### Option B: Import from Animations Package (Recommended)
Update `api/index.go` to import your animations package:

1. **Modify api/index.go**:
   ```go
   package handler

   import (
       "encoding/json"
       "fmt"
       "net/http"
       "strings"
       "time"
       "../animations" // Import your animations
   )

   const (
       clearScreen = "\033[2J\033[H"
   )

   func Handler(w http.ResponseWriter, r *http.Request) {
       // Use animations.SealWiggleFrames instead of local SealWiggleFrames
       frames := animations.SealWiggleFrames
       
       // ... rest of your handler code, but replace SealWiggleFrames with frames
   }
   ```

## Step 2: Configure Vercel

Your `vercel.json` is already set up correctly:

```json
{
  "functions": {
    "api/index.go": {
      "runtime": "@vercel/go"
    }
  },
  "rewrites": [
    {
      "source": "/silly_seal/(.*)",
      "destination": "/api/index.go?path=$1"
    },
    {
      "source": "/silly_seal",
      "destination": "/api/index.go"
    }
  ]
}
```

## Step 3: Deploy to Vercel

### Quick Deploy (Recommended)
```bash
# Login to Vercel (if not already)
vercel login

# Deploy from your project directory
cd /Users/rexliu/bs_ani_go
vercel --prod
```

### Manual Steps
1. **Initialize Vercel project**:
   ```bash
   vercel
   ```
   
2. **Answer the prompts**:
   - Set up and deploy? **Y**
   - Which scope? Choose your account
   - Link to existing project? **N** (if first time)
   - Project name? **bs-ani-go** or similar
   - Directory? **./** (current directory)
   - Auto-detected settings? **Y**

3. **Deploy to production**:
   ```bash
   vercel --prod
   ```

## Step 4: Test Your Deployment

After deployment, test your endpoints:

```bash
# Test the animation
curl https://your-deployment-url.vercel.app/silly_seal

# Test the API list
curl https://your-deployment-url.vercel.app/silly_seal/list

# Test in browser
open https://your-deployment-url.vercel.app/silly_seal
```

## Step 5: Custom Domain (Optional)

If you want to use a custom domain like `blubberstudios.com`:

1. **Add domain in Vercel dashboard**:
   - Go to your project settings
   - Click "Domains" 
   - Add your domain

2. **Update DNS**:
   - Add CNAME record pointing to `cname.vercel-dns.com`
   - Or add A record to Vercel's IP

3. **Update URLs in code**:
   - Update the URLs in `api/index.go` to use your custom domain

## Troubleshooting

### Large File Size Issue
Your 187-frame animation creates a ~3.5MB file. Vercel has limits:
- **Function size**: 50MB uncompressed
- **Response size**: 4.5MB for Hobby plan, 50MB for Pro

If you hit size limits:

1. **Compress frames** (remove extra whitespace):
   ```bash
   # This will reduce file size significantly
   sed -i 's/⠀//g' animations/seal_wiggle.go
   ```

2. **Reduce frame count** (use every 2nd frame):
   ```bash
   # Create a version with ~94 frames instead of 187
   ```

3. **Upgrade to Vercel Pro** for higher limits

### Import Path Issues
If you get import errors, ensure your `go.mod` is set up correctly:

```bash
# Check your go.mod file
cat go.mod

# Should contain:
# module seal-ascii
# go 1.21
```

### Function Timeout
Vercel serverless functions have timeout limits:
- **Hobby**: 10 seconds
- **Pro**: 15 seconds  
- **Enterprise**: 900 seconds

The current code limits animation to 30 seconds, which may exceed limits. Reduce in `api/index.go`:

```go
maxDuration := 10 * time.Second // Reduce for Hobby plan
```

## Alternative: Static Version

For a simpler deployment that avoids function limits, create a static version:

1. **Generate static HTML with embedded animation**:
   ```bash
   # Create a static version that cycles through frames in JavaScript
   ```

2. **Deploy as static site**:
   ```bash
   vercel --prod
   ```

## Files Checklist Before Deploy

Ensure these files are ready:
- ✅ `vercel.json` - Vercel configuration
- ✅ `api/index.go` - Updated with 187 frames  
- ✅ `animations/seal_wiggle.go` - Your 187 frames
- ✅ `go.mod` - Go module definition
- ✅ `go.sum` - Go dependencies

## Production URLs

After deployment, your endpoints will be:
- **Animation**: `https://your-app.vercel.app/silly_seal`
- **API List**: `https://your-app.vercel.app/silly_seal/list`
- **Browser**: `https://your-app.vercel.app/silly_seal` (shows instructions)

## Performance Tips

1. **Enable compression** in Vercel
2. **Add caching headers** for static content
3. **Optimize frame data** (remove unnecessary characters)
4. **Consider WebSocket** for real-time streaming instead of HTTP streaming

## Need Help?

- Check Vercel logs: `vercel logs`
- Vercel docs: [vercel.com/docs](https://vercel.com/docs)
- Test locally: `vercel dev`