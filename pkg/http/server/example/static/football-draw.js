// Champions League Draw Algorithm
// JavaScript version of the Go algorithm for football team draws

class FootballDraw {
    constructor() {
        this.teams = [];
        this.buckets = [[], [], [], []];
        this.matches = [];
    }

    // Add a team to the list
    addTeam(teamName) {
        if (teamName.trim() !== '') {
            this.teams.push(teamName.trim());
        }
    }

    // Remove a team by index
    removeTeam(index) {
        if (index >= 0 && index < this.teams.length) {
            this.teams.splice(index, 1);
        }
    }

    // Clear all teams
    clearTeams() {
        this.teams = [];
    }

    // Get current team count
    getTeamCount() {
        return this.teams.length;
    }

    // Check if we have exactly 36 teams
    isValidTeamCount() {
        return this.teams.length === 36;
    }

    // Shuffle array using Fisher-Yates algorithm
    shuffleArray(array) {
        const shuffled = [...array];
        for (let i = shuffled.length - 1; i > 0; i--) {
            const j = Math.floor(Math.random() * (i + 1));
            [shuffled[i], shuffled[j]] = [shuffled[j], shuffled[i]];
        }
        return shuffled;
    }

    // Generate the draw
    generateDraw() {
        if (!this.isValidTeamCount()) {
            throw new Error(`Expected 36 teams, got ${this.teams.length}`);
        }

        // Shuffle teams
        const shuffledTeams = this.shuffleArray(this.teams);

        // Distribute teams into 4 buckets of 9 teams each
        this.buckets = [[], [], [], []];
        for (let i = 0; i < 36; i++) {
            const bucketIndex = Math.floor(i / 9);
            const teamIndex = i % 9;
            this.buckets[bucketIndex][teamIndex] = shuffledTeams[i];
        }

        // Generate matches
        this.generateMatches();

        return {
            buckets: this.buckets,
            matches: this.matches
        };
    }

    // Generate all matches
    generateMatches() {
        this.matches = [];

        // Generate matches within each bucket
        for (const bucket of this.buckets) {
            this.matches.push(...this.generateBucketPairs(bucket));
        }

        // Generate cross-bucket matches
        const bucketPairs = [
            [this.buckets[0], this.buckets[1]],
            [this.buckets[0], this.buckets[2]],
            [this.buckets[0], this.buckets[3]],
            [this.buckets[1], this.buckets[2]],
            [this.buckets[1], this.buckets[3]],
            [this.buckets[2], this.buckets[3]]
        ];

        for (const [bucketA, bucketB] of bucketPairs) {
            this.matches.push(...this.generateCrossBucketPairs(bucketA, bucketB));
        }
    }

    // Generate matches within a bucket (3 groups of 3 teams each)
    generateBucketPairs(bucket) {
        const pairs = [];

        for (let i = 0; i < 3; i++) {
            const group = bucket.slice(i * 3, (i + 1) * 3);
            pairs.push(
                { teamA: group[0], teamB: group[1] },
                { teamA: group[2], teamB: group[0] },
                { teamA: group[1], teamB: group[2] }
            );
        }

        return pairs;
    }

    // Generate matches between two buckets
    generateCrossBucketPairs(bucketA, bucketB) {
        const pairs = [];

        for (let i = 0; i < bucketA.length; i++) {
            pairs.push({ teamA: bucketA[i], teamB: bucketB[i] });
            pairs.push({ teamA: bucketB[(i + 1) % bucketB.length], teamB: bucketA[i] });
        }

        return pairs;
    }

    // Format results for display
    formatResults() {
        const bucketNames = ['A', 'B', 'C', 'D'];
        let result = '';

        // Format groups
        for (let i = 0; i < this.buckets.length; i++) {
            result += `Group ${bucketNames[i]}:\n`;
            for (const team of this.buckets[i]) {
                result += `${team}\n`;
            }
            result += '\n';
        }

        // Format matches
        result += 'Matches:\n';
        for (const match of this.matches) {
            result += `${match.teamA} - ${match.teamB}\n`;
        }

        return result;
    }

    // Get groups as HTML
    getGroupsHTML() {
        const bucketNames = ['A', 'B', 'C', 'D'];
        let html = '';

        for (let i = 0; i < this.buckets.length; i++) {
            html += `<div class="group">`;
            html += `<h4>Group ${bucketNames[i]}</h4>`;
            html += `<ul>`;
            for (const team of this.buckets[i]) {
                html += `<li>${team}</li>`;
            }
            html += `</ul>`;
            html += `</div>`;
        }

        return html;
    }

    // Get matches as HTML
    getMatchesHTML() {
        let html = '<div class="matches-list">';
        for (const match of this.matches) {
            html += `<div class="match">${match.teamA} - ${match.teamB}</div>`;
        }
        html += '</div>';
        return html;
    }
}

// Export for use in HTML
window.FootballDraw = FootballDraw;
