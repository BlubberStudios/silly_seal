# 🚀 Quick Deployment Guide

## Ready to Deploy - Final Steps

Your seal animation is ready for Vercel deployment! Here's what you need to do:

### 1. Replace the API Handler (IMPORTANT)

You have two options:

**Option A: Use the optimized version I created:**
```bash
# Replace the current api/index.go with the optimized version
mv api/index.go api/index_old.go
mv api/index_new.go api/index.go
```

**Option B: Or manually copy all 24 frames:**
Copy the complete frames array from `vercel_frames_compact.txt` into `api/index.go` (this file has all 24 frames ready to use).

### 2. Deploy to Vercel

```bash
# Install Vercel CLI if you haven't already
npm i -g vercel

# Login to Vercel
vercel login

# Deploy from your project directory
cd /Users/rexliu/bs_ani_go
vercel --prod
```

### 3. Test Your Deployment

After deployment, Vercel will give you a URL like `https://bs-ani-go-xyz.vercel.app`

Test it:
```bash
# Test the animation
curl https://your-deployment-url.vercel.app/silly_seal

# Test in browser
open https://your-deployment-url.vercel.app/silly_seal
```

### 4. What You Get

- **24 smooth braille frames** (selected from your original 187 frames)
- **Vercel-optimized** for function size and timeout limits
- **Works with curl** for terminal animation
- **Browser-friendly** HTML page with instructions
- **API endpoint** at `/silly_seal/list`

### 5. Custom Domain (Optional)

If you want to use `blubberstudios.com/silly_seal`:

1. In Vercel dashboard → Your Project → Settings → Domains
2. Add your domain
3. Update DNS records as instructed by Vercel
4. Update the URLs in `api/index.go` to use your custom domain

### File Structure Summary

```
/Users/rexliu/bs_ani_go/
├── vercel.json           ✅ Already configured
├── api/index.go          🔄 Replace with optimized version
├── animations/           ✅ Your 187 frames (used locally)
├── go.mod               ✅ Ready
└── DEPLOYMENT.md        📖 Full documentation
```

### Deployment Commands (Copy & Paste)

```bash
# Quick deployment
cd /Users/rexliu/bs_ani_go
mv api/index.go api/index_old.go
mv api/index_new.go api/index.go
vercel --prod
```

That's it! Your 🦭 wiggling seal will be live on Vercel!

---

## Troubleshooting

**If deployment fails due to size:**
- The optimized version uses 24 frames (~440KB) which should work
- If still too large, reduce to 12 frames by changing the step from 8 to 16

**If function times out:**
- Reduce `maxDuration` in `api/index.go` from 15s to 10s
- Upgrade to Vercel Pro for longer timeouts

**Need help?**
- Check deployment logs: `vercel logs`
- Test locally: `vercel dev`