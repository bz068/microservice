const router = require('express').Router();
const { responseGenerator } = require('../utils');

const comments = [
    {
        id: 23,
        postId: 1,
        comment: 'first comment',
    },
];

router.post('/api/v1/comments', (req, res) => {
    try {
        const { postId, comment } = req.body;

        if (!postId) {
            res.json(responseGenerator('Please Provide a Post ID', false));
            return;
        }
        if (!comment) {
            res.json(responseGenerator('Please Provide a Comment', false));
            return;
        }

        const id = Math.floor(Math.random() * 601020);

        const newComment = { id, postId, comment };

        comments.push(newComment);

        res.json(
            responseGenerator('Comment Added Successfully', true, newComment)
        );
    } catch (error) {
        res.send('Internal Server Error');
    }
});

router.get('/api/v1/posts/:id/comments', (req, res) => {
    const id = req.params.id;

    const postComments = comments.filter(
        (comment) => comment.postId === parseInt(id)
    );

    if (postComments.length === 0) {
        res.json(responseGenerator('No Comments Found For This ID', false));
        return;
    }

    res.json(
        responseGenerator('Comment Fetched Successfully', true, postComments)
    );
});

// Routing Middleware
/* this middleware will be executed if route is not defined. */
router.use((req, res, next) => {
    res.json(
        responseGenerator('Route Not Found', false, [])
    );
    next()
  });

module.exports = router;
